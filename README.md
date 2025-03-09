
# Gooey

<p align="center">
    <img width="256" height="256" src="https://raw.githubusercontent.com/cookiengineer/gooey/master/assets/gooey.jpg">
</p>

Gooey (GUI) is a Pure Go Web UI framework made for stateless HTML in Web Views.
It bridges the gaps between Go, WebASM, Browser APIs, unified App Layouts and reusable Themes.


# Opinions

**HTML Elements**

- Static elements can never be removed from the DOM
- Static elements can have DOM event listeners
- Static elements always have an `id` property
- Dynamic elements can be removed from the DOM
- Dynamic elements should not have DOM event listeners

**App Architecture**

- [app.Main](/source/app/Main.go)
- [app.Client](/source/app/Client.go)
- [app.Storage](/source/app/Storage.go)
- [app.View](/source/app/View.go) interface
- [app.BaseView](/source/app/BaseView.go) class

**App Layout**

- App Layout always consists of `header`, `main`, and `footer` elements
- App Views are represented by different `main > section[data-view=...]` elements
- App Views can contain `aside` elements representing the sidebar

**Web Clients**

- Web Clients are REST API clients using either [fetch](/source/fetch) or [xhr](/source/xhr)
- Web Clients are fulfill an `interface` with `Validate() bool, err`
- Web Clients are routable via a `map[URL]struct` for each supported route

**Web Forms**

Note: This is currently work-in-progress

- Web Forms are static elements
- Web Forms with `enctype="application/json"` use a REST compatible API endpoint
- Web Forms with an `action` URL are automatically validateable
- Web Forms are serializable via `json.Marshal()` into a `struct`
- Web Forms are fulfill an `interface` with `Validate() bool, err`


## Implementation Details

- [docs/BINDINGS.md](/docs/BINDINGS.md) documents the state of implemented bindings
- [TODO.md](/TODO.md) documents the work-in-progress of things that will be implemented next


## Examples

The [examples](/examples) folder contains minimal test cases that show how you can
use the bindings. They also contain a separate `main.go` which is compiled into a
`main.wasm` file and a `serve.go` which reflects the local webserver.

All examples are served on `http://localhost:3000` if you execute the `build.sh`.

These examples also serve as unit tests, because `go test` cannot generate binaries
for the `syscall/js` platform.

- [crypto-aescbc](/examples/crypto-aescbc/main.go)
- [elements](/examples/elements/main.go)
- [fetch](/examples/fetch/main.go)
- [history](/examples/history/main.go)
- [location](/examples/location/main.go)
- [navigator](/examples/navigator/main.go)
- [navigator-geolocation](/examples/navigator-geolocation/main.go)
- [storages](/examples/storages/main.go)

## Projects

These are the Projects using `gooey` as a library. This list is meant to showcase how to use the
library, how to integrate it into your workflow, and how to integrate it with [webview/webview_go](https://github.com/webview/webview_go).

- [Git Evac](https://github.com/cookiengineer/git-evac), a Git Management Tool


# License

This project is licensed under the [MIT](./LICENSE_MIT.txt) license.

