
# Controller Implementation Guide

This guide explains how to implement a custom
[Controller](/components/app/Controller.go) for a Gooey App. It is the middle of the
[Reactive MVC Architecture](/docs/ARCHITECTURE.md): it receives events from a
[View](/docs/view-implementation-guide.md), talks to the Backend through `actions`, updates
the model, and triggers a render.

The companion guides are:

- [component-implementation-guide.md](/docs/component-implementation-guide.md)
- [view-implementation-guide.md](/docs/view-implementation-guide.md)
- [backend-implementation-guide.md](/docs/backend-implementation-guide.md)

**IMPORTANT**: Before reading this guide, read the [ERRATA.md](/docs/ERRATA.md) document.
It documents the quirks of running Go in a Web Browser and is relevant even for experienced
Go developers.

## Where a Controller sits in the Loop

```text
Controller ──action──▶ Backend
    ▲                     │
  event                update
    │                     ▼
  View ◀──render──── Model / Storage
```

1. The View fires an event (`action`, `change-value`, `change-field`, ...).
2. The Controller's event listener runs and maps the event to an `action`.
3. The `action` (in the `app/actions` package) calls the Backend via `app.Client`.
4. The Backend returns an `update` (a shared Schema).
5. The Controller writes the Schema to [app.Storage](/components/app/Storage.go) (the model).
6. The Controller updates its Components and calls `View.Render()`.

A Controller **must not** touch the DOM directly. It only talks to
[interfaces.Component](/components/interfaces/Component.go) methods, and lets the View own the
markup. A Controller **owns**:

- a reference to its `View`,
- a reference to `app.Main` (for `Client`, `Storage`, `Header`, `Footer`, `Dialog`),
- the current `Schema` (typed, deserialized from the Backend).

## Recommended App File Layout

The recommended frontend layout (see the
[View guide](/docs/view-implementation-guide.md#recommended-app-file-layout) for the full
tree) puts one Controller per View in `app/controllers`, plus a `RegisterTo.go`:

```text
app/
├── main.go
├── actions/
│   ├── GetTasks.go
│   ├── CreateTask.go
│   ├── UpdateTask.go
│   └── DeleteTask.go
├── components/
│   ├── RegisterTo.go
│   └── CustomComponent.go
├── controllers/
│   ├── RegisterTo.go
│   ├── Tasks.go
│   └── Settings.go
├── structs/
├── views/
│   ├── RegisterTo.go
│   └── Settings.go
└── public/
```

## The Controller Contract

A Controller implements [interfaces.Controller](/components/interfaces/Controller.go):

| Method        | Purpose                                                                       |
|:--------------|:------------------------------------------------------------------------------|
| `Name() string` | Lowercase identifier. Must match the View name and `data-name`.              |
| `Enter() bool`  | Called when the View becomes active. Wire event listeners and start `Update`. |
| `Leave() bool`  | Called when another View becomes active. Remove all event listeners.          |
| `Update()`      | Reads the model from the Backend and refreshes Components.                    |
| `Render()`      | Re-renders the View.                                                          |

**IMPORTANT**: As of Gooey `0.0.8`, [app.Main.ChangeView](/components/app/Main.go) calls both
the View's and the Controller's `Leave()` and `Enter()` methods correctly. Controllers are
created for every registered View during `app.Main.Mount()`, **after** the Component graph has
been built. You never call `Enter()`/`Leave()`/`Mount()` yourself.

## Implementing a Custom Controller

The walkthrough below implements the `Tasks` Controller from the
[app example](/examples/components/app/controllers/Tasks.go).

### 1. Struct and Constructor

The constructor receives `app.Main` and the already-instantiated `interfaces.View`. Type-assert
to your concrete View type. If the View is a *default* `app.View`, assert to `*app.View`; if it
is a *custom* View, assert to that type.

```go
//go:build wasm

package controllers

import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/interfaces"
import "example/schemas"

type Tasks struct {
	Main   *app.Main      `json:"main"`
	Schema *schemas.Tasks `json:"schema"`
	View   *app.View      `json:"view"`
}

func NewTasks(main *app.Main, view interfaces.View) *Tasks {

	var controller Tasks

	controller.Main   = main
	controller.Schema = &schemas.Tasks{}
	controller.View   = view.(*app.View)

	return &controller

}
```

The `Schema` field is the typed model. It is deserialized from the Backend response and is the
single source of truth for what the View renders.

### 2. Enter: Wire Event Listeners and Start Update

`Enter()` is where all event listeners are attached. Query the components from the View (and
from `Main.Dialog`/`Main.Footer`/`Main.Header`) and use
[components.UnwrapComponent](/components/UnwrapComponent.go) to recover the concrete type from
the interface returned by `Query()`.

```go
func (controller *Tasks) Enter() bool {

	// IMPORTANT: The Component Query API is self-including.

	fieldset, ok0 := components.UnwrapComponent[*content.Fieldset](controller.Main.Dialog.Query("dialog > fieldset"))
	table, ok1    := components.UnwrapComponent[*content.Table](controller.View.Query("section > article > table"))

	if fieldset != nil && table != nil && ok0 == true && ok1 == true {
		// ... attach listeners ...
	}

	go controller.Update()

	return true

}
```

Key facts:

- `Query()` returns `interfaces.Component`, hence `UnwrapComponent[*content.Table](...)`.
- `Query()` is self-including: `controller.View.Query("section > article > table")` starts at
  the View's own `<section>`.
- `Update()` is started in a goroutine so the UI does not block. It uses the `app.Client`,
  which is already asynchronous in the browser.

### 3. Event Listeners and `ToEventListener`

Add a listener with
[Component.AddEventListener](/components/Component.go) and wrap the callback with
[components.ToEventListener](/components/EventListener.go). The callback receives the event
name and a `map[string]any` of attributes.

```go
table.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

	action, ok := attributes["action"].(string)

	if ok == true {

		if action == "mark-done" {
			go func() {
				// ... run actions.UpdateTask(...) ...
			}()
		} else if action == "mark-undone" {
			go func() {
				// ... run actions.UpdateTask(...) ...
			}()
		}

	}

}, false))
```

The built-in events that you can listen for are:

| Event          | Fired by                                                                 |
|:---------------|:-------------------------------------------------------------------------|
| `click`        | [ui.Button](/components/ui/Button.go), Header, Footer, Dialog, Table.    |
| `change-value` | [ui.Input](/components/ui/Input.go), Checkbox, Number, Range, Select, Textarea. |
| `change-field` | [content.Fieldset](/components/content/Fieldset.go).                     |
| `change-view`  | [layout.Header](/components/layout/Header.go), [layout.Aside](/components/layout/Aside.go). |
| `action`       | Table, Footer, Dialog, custom Components.                                |

For `action` events, the meaningful attribute is `attributes["action"]`. For table rows, the
attributes map also contains the `data-*` attributes of the clicked element.

### 4. Footer, Dialog and Fieldset Patterns

Global UI lives on `app.Main` and is shared between Views. A common pattern:

```go
// Footer button opens the Dialog.
controller.Main.Footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

	action, ok := attributes["action"].(string)

	if ok == true && action == "create" {
		controller.Main.Dialog.Show()
	}

}, false))

// Dialog confirm reads the Fieldset and creates a Task.
controller.Main.Dialog.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

	action, ok := attributes["action"].(string)

	if ok == true {

		if action == "confirm" {

			go func() {

				title := fieldset.ValueOf("title").String()
				done  := fieldset.ValueOf("done").Bool()

				if title != "" {

					fieldset.Reset()
					controller.Main.Dialog.Disable()

					result, err := actions.CreateTask(controller.Main.Client, &schemas.Task{
						ID:    0,
						Title: title,
						Done:  done,
					})

					if err == nil && result.Title == title {
						table.Add(data.Data(map[string]any{
							"id":    result.ID,
							"title": result.Title,
							"done":  result.Done,
						}))
					}

					controller.Main.Dialog.Enable()
					controller.Main.Dialog.Hide()

				}

			}()

		} else if action == "cancel" {
			fieldset.Reset()
			controller.Main.Dialog.Enable()
			controller.Main.Dialog.Hide()
		}

	}

}, false))
```

Relevant methods:

- [layout.Dialog](/components/layout/Dialog.go): `Show()`, `Hide()`, `Enable()`, `Disable()`,
  `SetTitle()`, `SetContent()`.
- [layout.Footer](/components/layout/Footer.go): `Component.AddEventListener`, `Enable()`,
  `Disable()`, `SetContentLeft/Center/Right()`.
- [content.Fieldset](/components/content/Fieldset.go): `ValueOf(name) js.Value`,
  `Reset()`, `ResetField(name)`, `TypeOf(name)`.

### 5. Leave: Remove Event Listeners

`Leave()` must remove **every** listener that `Enter()` added. Passing `nil` removes all
listeners for that event, which is the recommended approach for shared UI such as the Footer
and Dialog, because another Controller re-attaches them on its own `Enter()`.

```go
func (controller *Tasks) Leave() bool {

	controller.Main.Footer.Component.RemoveEventListener("action", nil)
	controller.Main.Dialog.Component.RemoveEventListener("action", nil)

	return true

}
```

Also remove listeners from the View's own components. In the View guide we saw that a custom
View owns its children, so query them again in `Leave()` and detach:

```go
table, ok := components.UnwrapComponent[*content.Table](controller.View.Query("section > article > table"))

if table != nil && ok == true {
	table.Component.RemoveEventListener("action", nil)
}
```

### 6. Update: Action → Schema → Storage → Components

`Update()` performs the read part of the loop. It calls an `action`, stores the resulting Schema
in the Controller and in `Main.Storage`, updates the Components, and finally renders.

```go
func (controller *Tasks) Update() {

	if controller.Main != nil {

		schema, err := actions.GetTasks(controller.Main.Client)

		if err == nil {

			controller.Schema = schema
			controller.Main.Storage.Write("tasks", schema)

			table, ok := components.UnwrapComponent[*content.Table](controller.View.Query("section > article > table"))

			if ok == true && len(controller.Schema.Tasks) > 0 {

				dataset := data.NewDataset(0)

				for _, task := range controller.Schema.Tasks {
					dataset.Add(data.Data(map[string]any{
						"id":    task.ID,
						"title": task.Title,
						"done":  task.Done,
					}))
				}

				table.SetDataset(dataset)
				table.SortBy("id")

			}

		}

		controller.Render()

	}

}
```

[app.Storage](/components/app/Storage.go) is the model layer and uses the browser's
`LocalStorage`. Use a stable lowercase key (here `"tasks"`). `Storage.Read(name, schema)` is the
counterpart and is useful for optimistic rendering while `Update()` is still in flight.

### 7. Render

`Render()` always delegates to the View. The Controller is not allowed to manipulate the DOM.

```go
func (controller *Tasks) Render() {
	controller.View.Render()
}
```

### 8. Registration

Register the Controller with `app.Main` from `controllers/RegisterTo.go`:

```go
package controllers

import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/interfaces"

func RegisterTo(main *app.Main) {

	main.RegisterController("settings", func(main *app.Main, view interfaces.View) interfaces.Controller {
		return NewSettings(main, view)
	})

	main.RegisterController("tasks", func(main *app.Main, view interfaces.View) interfaces.Controller {
		return NewTasks(main, view)
	})

}
```

`app.WrapController` is the generic shorthand:

```go
func RegisterTo(main *app.Main) {
	main.RegisterController("tasks", app.WrapController(NewTasks))
}
```

**IMPORTANT**: The registered name must match the View's `data-name` and the
`main.RegisterView` name. `app.Main` creates one Controller per registered View during
`Mount()`, so a Controller without a matching View is never instantiated.

## The Actions Layer

`actions` functions are thin, typed wrappers around [app.Client](/components/app/Client.go).
One function per endpoint, one file per function. The Client methods are:

| Method                          | HTTP verb | Typical use        |
|:--------------------------------|:----------|:-------------------|
| `Read(path)`                    | `GET`     | List / read        |
| `Create(path, payload []byte)`  | `POST`    | Create             |
| `Update(path, payload []byte)`  | `PATCH`   | Partial update     |
| `Delete(path, payload []byte)`  | `DELETE`  | Remove             |

All paths must start with `/api`. The Client returns a `*fetch.Response` whose
`response.Body` is the raw JSON; `actions` unmarshal it into the shared Schema.

```go
package actions

import "github.com/cookiengineer/gooey/components/app"
import "example/schemas"
import "encoding/json"

func GetTasks(client *app.Client) (*schemas.Tasks, error) {

	var result_schema *schemas.Tasks = nil
	var result_error  error          = nil

	response, err1 := client.Read("/api/tasks")

	if err1 == nil {

		schema := schemas.Tasks{}
		err2   := json.Unmarshal(response.Body, &schema)

		if err2 == nil {
			result_schema = &schema
		} else {
			result_error = err2
		}

	} else {
		result_error = err1
	}

	return result_schema, result_error

}
```

**IMPORTANT**: Keep `actions` free of UI concerns. They must not import `app_views`,
`app_components`, or `app_controllers`. This keeps the dependency graph one-directional:
`Controllers → Actions → Client/Bindings`.

## Concurrency and Batch Actions

When an action applies to many items (a multi-select table), run the requests concurrently
with a `sync.WaitGroup`, and only render after all of them finished. The
[app example](/examples/components/app/controllers/Tasks.go) uses this pattern:

```go
indexes, dataset := table.Selected()
waitgroup := sync.WaitGroup{}

table.Deselect(indexes)

for d := 0; d < len(dataset); d++ {

	index := indexes[d]
	entry := dataset[d]

	waitgroup.Add(1)

	go func(index int, entry data.Data) {

		result, err := actions.UpdateTask(controller.Main.Client, &schemas.Task{
			ID:    entry["id"].(int),
			Title: entry["title"].(string),
			Done:  entry["done"].(bool),
		})

		if err == nil && result.Done == true {
			table.Remove([]int{index})
			table.Add(entry)
		}

		defer waitgroup.Done()

	}(index, entry)

}

waitgroup.Wait()
table.Render()
```

If you need progress feedback for long-running batches, keep a queue in `app/structs` and
render a dedicated progress Component as results arrive.

## Error Handling

The shipped example keeps error handling minimal, but real apps should not silently drop
errors. Recommended approach:

- Let `actions` return `(schema, error)` and check both the `error` and a sanity check on the
  returned Schema.
- On error, keep the previous `controller.Schema` and `Main.Storage` value so the UI stays
  consistent.
- Surface the error by setting a `ui.Label` label, or by rendering an error state in a custom
  Component.
- Never `panic` in an event listener: a panic in a WebASM goroutine terminates the whole
  program.

## Full Example

See
[examples/components/app/controllers/Tasks.go](/examples/components/app/controllers/Tasks.go)
for the complete, runnable file, and
[examples/components/app/controllers/Settings.go](/examples/components/app/controllers/Settings.go)
for the minimal skeleton.

## Controller Checklist

- [ ] File has the `//go:build wasm` build tag (if it references `app.View`/DOM types).
- [ ] Package is `controllers` and the file is named after the feature (`Tasks.go`).
- [ ] Struct has `Main`, `Schema`, and `View` fields.
- [ ] `NewTasks(main, view)` type-asserts the View once, in the constructor.
- [ ] `Name()` returns the same lowercase name as the View and `data-name`.
- [ ] `Enter()` attaches all listeners and calls `go controller.Update()`.
- [ ] `Leave()` removes every listener (`RemoveEventListener(event, nil)`).
- [ ] `Update()` calls an action, writes `Schema` to `Storage`, updates Components.
- [ ] `Render()` only delegates to `controller.View.Render()`.
- [ ] Actions are called through `controller.Main.Client`, never `fetch` directly.
- [ ] Registered in `controllers/RegisterTo.go` under the View's name.
- [ ] No DOM manipulation and no Backend knowledge inside the Controller.

## References

- [interfaces.Controller](/components/interfaces/Controller.go)
- [app.Controller](/components/app/Controller.go)
- [app.Main](/components/app/Main.go)
- [app.Client](/components/app/Client.go)
- [app.Storage](/components/app/Storage.go)
- [components.ToEventListener](/components/EventListener.go)
- [app.WrapController / app.UnwrapController](/components/app/WrapController.go)
- [components.WrapComponent / components.UnwrapComponent](/components/WrapComponent.go)
- [app example](/examples/components/app)
- [ARCHITECTURE.md](/docs/ARCHITECTURE.md)
