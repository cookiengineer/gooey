
# Gooey

<p align="center">
    <img width="256" height="256" src="https://raw.githubusercontent.com/cookiengineer/gooey/master/assets/gooey.jpg">
</p>

[Gooey](https://github.com/cookiengineer/gooey) (GUI) is divided in two parts:

- A pure Go WebASM bindings framework that bridges the gaps between Go, WebASM and Browser APIs.
- A pure Go Web UI Components framework that structures a Web Application, ready to be used in local Web Views.


## Gooey Bindings vs. Gooey Components

- The `github.com/cookiengineer/gooey/bindings` package contains all Web API Bindings.
- The `github.com/cookiengineer/gooey/components` package contains all Web UI/UX Components.


## Motivation

Problems in modern Web App Development:

- Components are great as a separation of feature concept.
- Components are bad for web accessibility (`aria-` property fatigue).
- Frontend-to-Backend communication is always a problem.
- Frontend Routers differ always from Backend Routers.
- Frontend schema safety and validation is always implemented redundantly.
- Backend schema safety and validation is great with `Marshal` / `Unmarshal`.

Conclusions:

- Let's use Go's types and schemas on the Frontend and the Backend.
- Let's use Gooey's Web Components for the Frontend via WebASM.
- Let's use Gooey's Web Components for the Backend via Go's native builds.
- Let's use Gooey to deploy local Apps via `webview/webview`.
- Let's use Gooey to deploy online Web Apps online.


## Architecture

Gooey uses a Reactive MVC Architecture and embraces the use of a unidirectional
flow, meaning it is a circular pattern of state management.

This is a nice architecture pattern for deserialization and serialization of the
Schemas into Views, because all custom application code lands in the controllers,
which manage the Views, Storages and networking code.

![Reactive MVC Architecture](/assets/reactive-mvc.jpg)


## Documentation

- [ARCHITECTURE.md](/docs/ARCHITECTURE.md) documents the architecture of a Gooey App.
- [BINDINGS.md](/docs/BINDINGS.md) documents the state of implemented web bindings.
- [COMPONENTS.md](/docs/COMPONENTS.md) documents the state of implemented web components.
- [TODO.md](/TODO.md) documents the work-in-progress of things that will be implemented in the near future.


## Examples

The [examples](/examples) folder contains minimal test cases that show how you can
use the bindings. They also contain a separate `main.go` which is compiled into a
`main.wasm` file and a `serve.go` which reflects the local webserver.

All examples are served on `http://localhost:3000` when the `build.sh` is executed.

Important: These examples also serve as unit tests, because `go test` cannot generate
binaries for the `syscall/js` platform. As soon as unit tests are available, the plan
is to migrate towards go-integrated tests that are compatible with `go test`.

**Bindings Examples**:

- [console](/examples/bindings/console)
- [crypto-aescbc](/examples/bindings/crypto-aescbc)
- [dom](/examples/bindings/dom)
- [fetch](/examples/bindings/fetch)
- [history](/examples/bindings/history)
- [location](/examples/bindings/location)
- [navigator](/examples/bindings/navigator)
- [navigator-geolocation](/examples/bindings/navigator-geolocation)
- [storages](/examples/bindings/storages)

**Components Examples**:

- [app](/examples/components/app)
- [content-fieldset](/examples/components/content-fieldset)
- [layout](/examples/components/layout)
- [ui](/examples/components/ui)


# License

This project is licensed under the [MIT](./LICENSE_MIT.txt) license.

