
# Gooey

<p align="center">
    <img width="256" height="256" src="https://raw.githubusercontent.com/cookiengineer/gooey/master/assets/gooey.jpg">
</p>

Gooey (GUI) is a Pure Go WebASM bindings framework. It bridges the gaps between Go, WebASM and Browser APIs.

On top of Gooey, there's the [Gooey Components](https://github.com/cookiengineer/gooey-components) framework,
which offers ready-to-use Web Components to structure a Web Application that uses a local Web View for its UI.


## Bindings

- [docs/BINDINGS.md](/docs/BINDINGS.md) documents the state of implemented bindings.
- [TODO.md](/TODO.md) documents the work-in-progress of things that will be implemented in the near future.


## Examples

The [examples](/examples) folder contains minimal test cases that show how you can
use the bindings. They also contain a separate `main.go` which is compiled into a
`main.wasm` file and a `serve.go` which reflects the local webserver.

All examples are served on `http://localhost:3000` if you execute the `build.sh`.

These examples also serve as unit tests, because `go test` cannot generate binaries
for the `syscall/js` platform.

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

