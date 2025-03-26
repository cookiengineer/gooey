
# Gooey

<p align="center">
    <img width="256" height="256" src="https://raw.githubusercontent.com/cookiengineer/gooey/master/assets/gooey.jpg">
</p>

[Gooey](https://github.com/cookiengineer/gooey) (GUI) is divided in two parts:

- A pure Go WebASM bindings framework that bridges the gaps between Go, WebASM and Browser APIs.
- A pure Go Web UI Components framework that structures a Web Application, ready to be used in local Web Views.


## Bindings vs. Components

- The `github.com/cookiengineer/gooey/bindings` package contains all Web API Bindings.
- The `github.com/cookiengineer/gooey/components` package contains all Web UI/UX Components.

For further details, take a look at the following files to get an overview:

- [BINDINGS.md](/docs/BINDINGS.md) documents the state of implemented web bindings.
- [COMPONENTS.md](/docs/COMPONENTS.md) documents the state of implemented web components.
- [TODO.md](/TODO.md) documents the work-in-progress of things that will be implemented in the near future.


## Program Architecture

- Program serves a local Web UI and opens a webview pointing towards the UI.
- Program uses `go:embed` to embed a `/public/*` folder that contains all assets.
- Program uses a UI using `HTML`, `CSS`, and `WebASM`.
- Program uses a Reactive MVC architecture, which allows a circular flow.

![Reactive MVC Architecture](/assets/reactive-mvc.jpg)

## WebView Architecture

The HTML/DOM App Layout always consists of the following elements:

- `body > header`, represented by [layout/Header](/components/layout/Header.go)
- `body > main`, represented by [app/Main](/components/app/Main.go)
- `body > main > section[data-view=...]`, represented by [app/View](/components/app/View.go)
- `body > main > section[data-view=...] > aside`, represented by [layout/Aside](/components/layout/Aside.go)
- `body > footer`, represented by [layout/Footer](/components/layout/Footer.go)

The CSS App Theme always uses semantic HTML as style rules:

- Activity/Visibility states of elements are represented by the `class="active"` attribute.
- All other element states are represented using `data-` attributes.

The [/design](/design) folder contains a ready-to-use default theme. Just copy the design folder so
that it is located at `/public/design/` and reachable as `/design/index.css`. Include the
`<link rel="stylesheet" href="/design/index.css">` in your HTML files.


## Examples

The [examples](/examples) folder contains minimal test cases that show how you can
use the bindings. They also contain a separate `main.go` which is compiled into a
`main.wasm` file and a `serve.go` which reflects the local webserver.

All examples are served on `http://localhost:3000` when the `build.sh` is executed.

Important: These examples also serve as unit tests, because `go test` cannot generate
binaries for the `syscall/js` platform. As soon as unit tests are available, the plan
is to migrate towards go-integrated tests that are compatible with `go test`.

- [console](/examples/console)
- [crypto-aescbc](/examples/crypto-aescbc)
- [elements](/examples/elements)
- [fetch](/examples/fetch)
- [history](/examples/history)
- [location](/examples/location)
- [navigator](/examples/navigator)
- [navigator-geolocation](/examples/navigator-geolocation)
- [storages](/examples/storages)


# License

This project is licensed under the [MIT](./LICENSE_MIT.txt) license.

