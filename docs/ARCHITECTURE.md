
# Architecture

![Reactive MVC Architecture](/assets/reactive-mvc.jpg)

## Program Architecture

- Program serves a local Web UI and opens a webview pointing towards the UI.
- Program uses `go:embed` to embed a `/public/*` folder that contains all assets.
- Program uses a UI using `HTML`, `CSS`, and `WebASM`.
- Program uses a Reactive MVC architecture, which allows a circular flow.
- Program uses on the frontend and backend side the same `struct` / [interfaces.Schema](/interfaces/Schema.go).

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

