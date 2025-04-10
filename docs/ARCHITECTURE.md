
# Architecture

![Reactive MVC Architecture](/assets/reactive-mvc.jpg)

## Program Architecture

- Program serves a local Web UI and opens a webview pointing towards the UI.
- Program uses `go:embed` to embed a `/public/*` folder that contains all assets.
- Program uses a UI using `HTML`, `CSS`, and `WebASM`.
- Program uses a Reactive MVC architecture, which allows a circular flow.
- Program uses on the frontend and backend side the same `struct` / [interfaces.Schema](/interfaces/Schema.go).

## Frontend and Backend

By default, Gooey embraces the use of a [REST Client](/components/app/Client.go)
via the [Fetch API](/bindings/fetch) bindings. You can also integrate your own
GraphQL client, given that it runs on the `syscall/js` platform.

Note: At a later point, gooey wants to offer a centralized Frontend and Backend Router
which can map the routes correctly to the views and REST APIs automatically.

## WebView HTML/DOM

The HTML/DOM App Layout always consists of the following elements:

- `body > header`, represented by [layout/Header](/components/layout/Header.go)
- `body > main`, represented by [app/Main](/components/app/Main.go)
- `body > main > section[data-view=...]`, represented by [app/View](/components/app/View.go)
- `body > main > section[data-view=...] > aside`, represented by [layout/Aside](/components/layout/Aside.go)
- `body > footer`, represented by [layout/Footer](/components/layout/Footer.go)

## WebView Theme

The CSS App Theme always uses semantic HTML as style rules:

- UI Components are native to be Web Accessibility compatible.
- Gooey CSS Theme is classless and instead based on `data-` attributes.
- User-provided Themes can use CSS classes for better customization.
- `data-action` influences the event flow (see Reactive MVC architecture).
- `data-layout` influences the layout (`grid`, `flex`, `flow`)
- `data-state` influences the activity/visibility (`active`)

The [/design](/design) folder contains a ready-to-use default theme.
Copy it, modify it, and you're good to go.

