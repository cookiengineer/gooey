
# Backend Implementation Guide

This guide explains how to implement the Backend of a Gooey App. A Gooey App is a local,
offline-first Web App: the Backend serves the WebASM frontend and a local JSON API, and it
shares the same Schema definitions with the frontend. It is the right-hand side of the
[Reactive MVC Architecture](/docs/ARCHITECTURE.md):

```text
Controller ──action──▶ Backend
    ▲                     │
  event                update
    │                     ▼
  View ◀──render──── Model / Storage
```

The companion guides are:

- [component-implementation-guide.md](/docs/component-implementation-guide.md)
- [view-implementation-guide.md](/docs/view-implementation-guide.md)
- [controller-implementation-guide.md](/docs/controller-implementation-guide.md)

**IMPORTANT**: Before reading this guide, read the [ERRATA.md](/docs/ERRATA.md) document.
It documents the quirks of running Go in a Web Browser and is relevant even for experienced
Go developers.

## Responsibilities of the Backend

The Backend **must**:

- serve the static WebASM assets (`main.wasm`, `wasm_exec.js`, HTML, CSS),
- expose a local JSON API under `/api/*`,
- hold the authoritative application state (the model),
- provide the `update` (a shared Schema) for every `action`,
- share Schema definitions with the frontend so that `Marshal`/`Unmarshal` round-trips are
  type-safe on both sides.

The Backend **must not**:

- render HTML for the SPA (the View layer owns rendering),
- contain UI logic or DOM bindings,
- trust any frontend input without validation.

## Recommended Backend File Layout

The frontend lives in `/app` and the Backend in `/source`. They are two Go modules that share
the `schemas` package of the Backend module via a `replace` directive in the App module.

```text
source/                          # native backend module
├── go.mod                       # module example-backend
├── cmds/
│   ├── example/main.go          # production binary, embeds public/
│   └── example-debug/main.go    # development server with hot reload
├── server/
│   ├── Dispatch.go              # static asset server + global endpoints
│   ├── DispatchRoutes.go        # the /api/* route table
│   ├── DispatchHotReload.go     # dev-only rebuild of main.wasm
│   ├── Serve.go                 # http.ListenAndServe wrapper
│   └── routes/
│       ├── Tasks.go             # GET /api/tasks, POST /api/tasks
│       └── Task.go              # GET/PATCH/DELETE /api/tasks/{id}
├── actions/                     # domain operations (no HTTP, no JSON)
│   ├── GetTasks.go
│   ├── CreateTask.go
│   ├── UpdateTask.go
│   └── DeleteTask.go
├── schemas/                     # shared wire format (imported by /app)
│   ├── Task.go
│   ├── Tasks.go
│   └── Settings.go
├── structs/                     # in-memory model / registry
│   └── Profile.go
├── types/                       # domain types with build-tagged methods
│   ├── Task.go
│   ├── Task_wasm.go
│   └── Task_other.go
├── services/                    # external API clients (optional)
├── utils/                       # pure helpers
└── public/                      # embedded static assets (built into /app/public)
    └── FS.go                    # //go:embed
```

The package responsibilities are:

| Folder             | Responsibility                                                             |
|:-------------------|:---------------------------------------------------------------------------|
| `cmds`             | `main()` entrypoints. Production vs. debug wiring.                         |
| `server`           | HTTP plumbing: static files, routing, hot reload.                          |
| `server/routes`    | One file per resource. Verb handling, JSON, status codes.                  |
| `actions`          | Domain logic. Validates input, mutates the model, returns typed errors.    |
| `schemas`          | The wire format, shared verbatim with the frontend.                        |
| `structs`          | In-memory state. Thread-safe registries and `Snapshot*` methods.           |
| `types`            | Domain objects. May have platform-specific (`_wasm` / `_other`) methods.   |
| `services`         | Clients for external APIs (optional, see the [services pattern](#services)). |
| `utils`            | Pure helpers (validation, string/path handling).                           |
| `public`           | Static assets and the App binary, embedded with `go:embed`.                |

The App module declares the Backend module as a replacement, so it can import the shared
schemas:

```go
// app/go.mod
module example-app

go 1.24.0

replace example-backend => ../source

require example-backend v0.0.0
require github.com/cookiengineer/gooey v0.0.0
```

```go
// source/go.mod
module example-backend

go 1.24.0
```

## Shared Schemas

The Schemas are the contract between View/Controller and Backend. Define them once, as small
structs with JSON tags, and import them from both sides. Never duplicate a Schema in the App
module.

```go
package schemas

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
```

Collection responses use an envelope so that they can grow without breaking clients:

```go
package schemas

type Tasks struct {
	Tasks map[int]*Task `json:"tasks"`
}
```

Single-resource responses return the resource directly (no envelope). This keeps the frontend
`actions` simple: `json.Unmarshal(response.Body, &schemas.Task{})` or
`&schemas.Tasks{}`.

Add validation methods to the Schema when the Backend must reject invalid input. For example,
`func (schema *Settings) IsValid() bool` can verify required fields, ranges, and enumerations
before persisting.

**IMPORTANT**: Because the same struct is compiled into the `js/wasm` frontend and the native
Backend, avoid build tags in `schemas`. If a Schema needs a platform-specific method, put it in
`types` and use a build tag (see [Platform-Specific Types](#platform-specific-types)).

## HTTP and JSON Conventions

### Paths

Use the Go 1.22+ standard library mux patterns. A resource has a collection route and an item
route:

```text
GET    /api/tasks          # list
POST   /api/tasks          # create
GET    /api/tasks/{id}     # read one
PATCH  /api/tasks/{id}     # partial update
PUT    /api/tasks/{id}     # full replace
DELETE /api/tasks/{id}     # remove
```

Read path variables with `request.PathValue("id")`. Always parse and validate them before use.

### Verbs and Idempotency

Choose the verb that reflects the idempotency of the operation. The frontend
[app.Client](/components/app/Client.go) maps directly to these verbs:

| Verb     | Method             | Idempotent | Use for                          |
|:---------|:-------------------|:-----------|:---------------------------------|
| `GET`    | `client.Read`      | yes        | Reading/list. No side effects.   |
| `POST`   | `client.Create`    | no         | Creating a new resource.         |
| `PATCH`  | `client.Update`    | yes        | Partial update.                  |
| `PUT`    | (raw fetch)        | yes        | Full replacement.                |
| `DELETE` | `client.Delete`    | yes        | Removal.                         |

### Responses

- Always set `Content-Type: application/json`.
- Use accurate status codes: `200` for success, `400` for malformed input, `404` for a missing
  resource, `405` for a wrong verb, `409` for conflicts, `500` for internal errors.
- On error, return an empty object `{}` for item endpoints and an empty array `[]` for
  collection endpoints. This keeps `json.Unmarshal` from failing on the frontend.
- Log each request as `> METHOD /api/path: <status text>`.

## Minimal Backend

The simplest possible Backend is a single `serve.go` that serves `public/` and implements the
JSON API. This is enough for the [app example](/examples/components/app/serve.go):

```go
package main

import "example-backend/schemas"
import "encoding/json"
import "log"
import "net/http"

func main() {

	fsrv := http.FileServer(http.Dir("public"))
	tasks := make(map[int]*schemas.Task)

	tasks[1] = &schemas.Task{ID: 1, Title: "Check out Gooey", Done: true}

	http.Handle("/", fsrv)

	http.HandleFunc("/api/tasks", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			payload, err := json.MarshalIndent(schemas.Tasks{Tasks: tasks}, "", "\t")

			if err == nil {

				log.Println("> GET /api/tasks: ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {
				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte("[]"))
			}

		} else {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("[]"))
		}

	})

	log.Println("Listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
```

For anything beyond a demo, split this into the layered structure below so that the HTTP layer,
the domain logic, and the model stay independently testable.

## Layered Backend

### Routes

A route file deals only with HTTP: verb guard, path variables, calling an action, refreshing
the model, marshalling the Schema, and status codes. No business logic.

```go
package routes

import "example-backend/actions"
import "example-backend/schemas"
import "example-backend/structs"
import "encoding/json"
import "io"
import "net/http"
import "strconv"

func Task(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	id, err := strconv.Atoi(request.PathValue("id"))

	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{}"))
		return
	}

	switch request.Method {

	case http.MethodGet:

		task := profile.GetTask(id)

		if task == nil {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte("{}"))
			return
		}

		payload, _ := json.MarshalIndent(task, "", "\t")

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(payload)

	case http.MethodPatch:

		bytes, _ := io.ReadAll(request.Body)

		patch := schemas.Task{}

		// TODO: json.Unmarshal(bytes, &patch) with error handling.

		task, err := actions.UpdateTask(profile, id, &patch)

		if err != nil {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte("{}"))
			return
		}

		payload, _ := json.MarshalIndent(task, "", "\t")

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(payload)

	default:

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{}"))

	}

}
```

### Actions

Actions contain the domain logic. They validate input, resolve entities from the model, perform
side effects (filesystem, `exec.Command`, external APIs), and return typed errors. They never
import `net/http` and never write JSON.

```go
package actions

import "example-backend/structs"
import "example-backend/schemas"
import "errors"

func UpdateTask(profile *structs.Profile, id int, patch *schemas.Task) (*schemas.Task, error) {

	task := profile.GetTask(id)

	if task == nil {
		return nil, errors.New("Task does not exist")
	}

	if patch.Title != "" {
		task.Title = patch.Title
	}

	task.Done = patch.Done

	return task, nil

}
```

### Model

Keep the in-memory model in `structs`. Protect every map with a `sync.RWMutex`, and expose
copying `Snapshot*()` accessors so that HTTP handlers cannot race with a refresh loop.

```go
package structs

import "example-backend/schemas"
import "io/fs"
import "sync"

type Profile struct {
	mutex      sync.RWMutex
	Tasks      map[int]*schemas.Task
	Filesystem *fs.FS
}

func NewProfile() *Profile {

	var profile Profile

	profile.Tasks = make(map[int]*schemas.Task)

	return &profile

}

func (profile *Profile) SnapshotTasks() map[int]*schemas.Task {

	profile.mutex.RLock()
	defer profile.mutex.RUnlock()

	result := make(map[int]*schemas.Task, len(profile.Tasks))

	for id, task := range profile.Tasks {
		result[id] = task
	}

	return result

}

func (profile *Profile) GetTask(id int) *schemas.Task {

	profile.mutex.RLock()
	defer profile.mutex.RUnlock()

	return profile.Tasks[id]

}
```

### Dispatch and Routing

[components.Document](/components/Document.go) is the frontend counterpart of this section; on
the Backend, `server.Dispatch` serves static assets and global endpoints, and
`server.DispatchRoutes` registers the `/api/*` route table.

```go
package server

import "example-backend/server/routes"
import "example-backend/structs"
import "net/http"

func DispatchRoutes(profile *structs.Profile) bool {

	http.HandleFunc("/api/tasks", func(response http.ResponseWriter, request *http.Request) {
		routes.Tasks(profile, request, response)
	})

	http.HandleFunc("/api/tasks/{id}", func(response http.ResponseWriter, request *http.Request) {
		routes.Task(profile, request, response)
	})

	return true

}
```

### Serve and main

`cmds/example/main.go` wires the model, dispatches, serves, and opens the WebView. Use
`go:embed` for production and `os.DirFS` for development.

```go
package main

import "example-backend/public"
import "example-backend/server"
import "example-backend/structs"
import "io/fs"

func main() {

	fsys, _ := fs.Sub(public.FS, ".")

	profile := structs.NewProfile()
	profile.Filesystem = &fsys
	profile.Refresh()

	server.Dispatch(profile)
	server.DispatchRoutes(profile)

	// Optionally open a WebView pointing at http://localhost:3000/index.html
	// with a webview binding of your choice.

	server.Serve(profile)

}
```

Always bind to `localhost` and a configurable port. A Gooey App is a local application, not a
public web service. The WebView itself is provided by an external package such as
`webview/webview`; Gooey does not ship a WebView binding.

## Embedding Static Assets

Use `go:embed` so the production binary is self-contained. Place the marker next to the assets:

```go
package public

import "embed"

//go:embed all:.
var FS embed.FS
```

Serve the embedded filesystem with `http.FS`:

```go
fsys := http.FS(public.FS)
fsrv := http.FileServer(fsys)
http.Handle("/", fsrv)
```

For the App to run in a browser, `index.html` must include the CSP directives that permit
WebASM evaluation:

```html
<meta http-equiv="Content-Security-Policy" content="default-src 'self' 'unsafe-eval' 'wasm-unsafe-eval'">
<meta http-equiv="Content-Security-Policy" content="script-src 'self' 'unsafe-eval' 'wasm-unsafe-eval'">
<meta http-equiv="Content-Security-Policy" content="script-src-elem 'self'">
<meta http-equiv="Content-Security-Policy" content="frame-src 'self'">
<meta http-equiv="Content-Security-Policy" content="worker-src 'self'">
<meta http-equiv="Content-Security-Policy" content="connect-src 'self'">
```

The `wasm-unsafe-eval` directive is required because WebASM's `JSON.parse`/`JSON.stringify`
integration compiles at runtime.

## Hot Reload (Development)

The debug entrypoint (`cmds/example-debug/main.go`) serves `public/` from disk
(`os.DirFS("public")`) and registers `server.DispatchHotReload`. That handler rebuilds
`main.wasm` from `/app` and `wasm_exec.js` from `$GOROOT` on every `GET /main.wasm` and
`GET /wasm_exec.js`, so developers can edit Go code and just reload the browser.

```go
fsys := os.DirFS("public")
profile := structs.NewProfile()
profile.Filesystem = &fsys

server.Dispatch(profile)
server.DispatchRoutes(profile)
server.DispatchHotReload(profile)
server.Serve(profile)
```

Set `Cache-Control: no-store` on the WebASM responses so the browser always fetches the fresh
binary. `server.DispatchHotReload` is application code that lives in your App's `server`
package; it is not part of Gooey.

## Build Tags for Platform-Specific Types

Some domain types need real methods on the Backend, but no-ops on the frontend (where there is
no filesystem or process access). Split them by platform:

```go
// types/Task.go
package types

type Task struct {
	Name   string
	Status bool
}

func NewTask(name string) *Task {
	return &Task{Name: name}
}
```

```go
// types/Task_other.go
//go:build !wasm

package types

func (task *Task) Refresh() bool {
	// Read from disk, call an external tool, etc.
	return true
}
```

```go
// types/Task_wasm.go
//go:build wasm

package types

func (task *Task) Refresh() bool {
	// Frontend stub: there is no filesystem in the browser.
	return false
}
```

The App still compiles the domain types, but only the `_wasm` variant is ever linked.

## Services

If the Backend talks to external APIs (for example a Git forge), isolate each provider in
`services/<provider>/`. A service package:

- takes a base URL, an owner/namespace, and a token,
- performs the HTTP requests and unmarshals into `types`,
- returns plain domain objects, not HTTP responses,
- never touches the model or the routes.

Dispatch to the correct provider from the model refresh code, based on a service type stored in
the Settings. A `types.Service` value with a `Name`, `URL`, `Type`, and `Token` field is a
practical shape for this, and the `Type` field selects which `services/<provider>` package to
call.

## Security

A Gooey Backend is local, but it still handles untrusted input from a browser context.

- **Validate names and IDs.** Reject values containing `/`, `\`, `..`, control characters, or
  leading `.` before using them in paths.
- **Constrain filesystem access.** Build paths with `filepath.Join(root, name)` and verify the
  result stays inside `root` (`filepath.Clean` + `strings.HasPrefix`).
- **Never build shell strings.** Use `exec.Command(binary, arg1, arg2, ...)` so arguments
  cannot be interpreted by a shell.
- **Do not log secrets.** Tokens and SSH keys must never end up in the console or in JSON.
- **Keep the CSP strict.** Only allow the directives WebASM actually needs.
- **Bind to localhost.** Do not expose the API on the LAN unless the App is explicitly designed
  for it.
- **Return generic error bodies.** `{}`/`[]` leak no internal state; log details server-side.

## Adding an Endpoint End to End

Every new feature touches both modules. Work through this checklist:

- [ ] **Schema** — add/extend the struct in `source/schemas` (shared with `/app`).
- [ ] **Action** — implement the domain operation in `source/actions`.
- [ ] **Route** — add verb handling and JSON in `source/server/routes`.
- [ ] **Dispatch** — register the path in `source/server/DispatchRoutes.go`.
- [ ] **Frontend action** — add a thin wrapper in `app/actions` that uses `app.Client`.
- [ ] **Controller** — call the action, write the Schema to `Storage`, update Components.
- [ ] **View** — add the `data-action` button/element and the `data-*` markup.
- [ ] **Test** — exercise the endpoint with `curl` before wiring the frontend.
- [ ] **Idempotency** — double-check the verb matches the semantics.

## Full Example

See [examples/components/app/serve.go](/examples/components/app/serve.go) for a complete,
runnable Backend that shares its Schemas with the WebASM frontend, including `GET`/`POST`
`/api/tasks` and `GET`/`PATCH` `/api/tasks/{id}`.

## References

- [examples/components/app](/examples/components/app)
- [app.Client](/components/app/Client.go)
- [app.Storage](/components/app/Storage.go)
- [ARCHITECTURE.md](/docs/ARCHITECTURE.md)
- [ERRATA.md](/docs/ERRATA.md)
- [view-implementation-guide.md](/docs/view-implementation-guide.md)
- [controller-implementation-guide.md](/docs/controller-implementation-guide.md)
