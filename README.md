
# Gooey

<p align="center">
    <img width="256" height="256" src="https://raw.githubusercontent.com/cookiengineer/gooey/master/assets/gooey.jpg">
</p>

[Gooey](https://github.com/cookiengineer/gooey) (pronounced as `/'ɡu.i/` or `GUI`) is divided in two parts:

- A pure Go WebASM [bindings](/bindings) framework that bridges the gaps between Go, WebASM and Browser APIs.
- A pure Go Web UI [components](/components) framework that structures a Web Application, ready to be used in local Web Views.


## Motivation

Problems in modern Web App Development:

- Web Components are great as a separation of feature concept.
- Web Components are bad for web accessibility (`aria-` property fatigue).
- Frontend-to-Backend communication is always a problem.
- Frontend schema safety and validation is always implemented redundantly in another language (be it ECMAScript, TypeScript, or whatever).
- Backend schema safety and validation is great with `Marshal` / `Unmarshal`, but is hard to keep on bug-to-bug parity with Frontend.
- Using online-first Web Apps with slow internet connections is very painful.

Conclusions:

- Use Go's types, structs and schemas on the Frontend via WebASM and on the Backend via its native builds.
- Use dynamic Web Components for the Frontend via WebASM.
- Use static Web Components for the Backend via Go's native builds to provide server-side rendering.
- Deploy offline-first Apps via `webview/webview` that point towards a local web server.
- Bundle all assets in `/public` via `go:embed` with the application binary.


## Architecture

Gooey uses a Reactive MVC Architecture and embraces the use of a unidirectional
flow, meaning it is a circular pattern of state management.

This is a nice architecture pattern for deserialization and serialization of the
Schemas into Views, because all custom application code lands in the separated
controllers, which each manage their own View, Client, and Storage.

![Reactive MVC Architecture](/assets/reactive-mvc.jpg)

Custom Controllers with even Custom Views are easily integratable this way, because
the Reactive MVC Architecture allows to build an App where each of the Views is
just a Web Component layer and doesn't need to be touched for most cases.


## Documentation

- [ERRATA.md](/docs/ERRATA.md) documents the state of known errata and problems of using Go via WebASM.

**IMPORTANT**: Note that even if you have years of Go development experience, the Errata
document is still relevant for you, because it highlights problems when using Go in the
Web Browser and the quirks that come with it.

- [ARCHITECTURE.md](/docs/ARCHITECTURE.md) documents the architecture of a Gooey App.
- [BINDINGS.md](/docs/BINDINGS.md) documents the state of implemented Web Bindings.
- [COMPONENTS.md](/docs/COMPONENTS.md) documents the state of implemented Web Components.
- [TODO.md](/docs/TODO.md) documents the work-in-progress of things that will be implemented in the near future.


## Examples

The [examples](/examples) folder contains minimal test cases that show how you can
use the bindings. They also contain a separate `main.go` which is compiled into a
`main.wasm` file and a `serve.go` which reflects the local webserver.

All examples are served on `http://localhost:3000` when the `build.sh` is executed.

**Important**: The examples also serve as unit tests, because `go test` cannot generate
binaries for the `syscall/js` platform right now. As soon as unit tests are available
upstream via `go test -c`, the plan is to migrate towards fully integrated unit tests.

**Bindings Examples**:

- [canvas2d](/examples/bindings/canvas2d)
- [console](/examples/bindings/console)
- [crypto-aescbc](/examples/bindings/crypto-aescbc)
- [cookiestore](/examples/bindings/cookiestore)
- [dom](/examples/bindings/dom)
- [encoding](/examples/bindings/encoding)
- [fetch](/examples/bindings/fetch)
- [history](/examples/bindings/history)
- [location](/examples/bindings/location)
- [navigator](/examples/bindings/navigator)
- [navigator-geolocation](/examples/bindings/navigator-geolocation)
- [storages](/examples/bindings/storages)

**Components Examples**:

- [app](/examples/components/app)
- [app-components](/examples/components/app-components)
- [content-fieldset](/examples/components/content-fieldset)
- [content-linechart](/examples/components/content-linechart)
- [content-piechart](/examples/components/content-piechart)
- [content-table](/examples/components/content-table)
- [layout](/examples/components/layout)
- [ui](/examples/components/ui)


# License

This project is licensed under the [MIT](./LICENSE.txt) license.


# Contributors

- [cookiengineer](https://github.com/cookiengineer)
- [utsavanand2](https://github.com/utsavanand2)

