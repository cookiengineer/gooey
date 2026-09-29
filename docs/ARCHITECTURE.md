
# Architecture

Gooey uses a Reactive MVC Architecture and embraces the use of a unidirectional
flow, meaning it is a circular pattern of state management.

This is a nice architecture pattern for deserialization and serialization of the
Schemas into Views, because all custom application code lands in the separated
controllers, which each manage their own View, Client, and Storage.

![Reactive MVC Architecture](/assets/reactive-mvc.jpg)

Custom Controllers with even Custom Views are easily integratable this way, because
the Reactive MVC Architecture allows to build an App where each of the Views is
just a Web Component layer and doesn't need to be touched for most cases.

## Program Architecture

- Program serves a local webserver and opens a `webview/webview` pointing towards the Web UI.
- Program uses `go:embed` to embed a `/public/*` folder that contains all assets.
- Program uses a Web UI using `HTML`, `CSS`, and `WebASM`.
- Program uses a Reactive MVC Architecture, which allows a circular flow of data.
- Program uses on the Frontend and Backend the same `Schema` definitions.

## Network Connections

By default, Gooey embraces the use of a [REST Client](/components/app/Client.go)
via the [Fetch API](/bindings/fetch) bindings. You can also integrate your own
GraphQL client, given that it runs on the `syscall/js` platform.

Note: At a later point, gooey wants to offer a centralized Frontend and Backend Router
which can map the routes correctly to the views and REST APIs automatically.

## WebView HTML/DOM

The HTML/DOM App Layout always consists of the following elements:

- `body > header`, represented by [layout/Header](/components/layout/Header.go)
- `body > main`, represented by [app/Main](/components/app/Main.go)
- `body > main > section[data-name=...]`, represented by [app/View](/components/app/View.go)
- `body > aside`, represented by [layout/Aside](/components/layout/Aside.go)
- `body > footer`, represented by [layout/Footer](/components/layout/Footer.go)
- `body > dialog`, represented by [layout/Dialog](/components/layout/Dialog.go)

`header` and `aside` are equivalent navigation containers and can be used interchangeably;
`app.Main` registers each View with any of them that is present in the markup.

Therefore minimal semantically enforced HTML5 code looks like this:

```html
<!DOCTYPE html>
<html>
	<head>(...)</head>
	<body>
		<!-- app.Main.Header -->
		<header>
			<div></div>
			<ul>
				<li data-state="active"><a data-view="welcome" href="/index.html">Welcome</a></li>
				<li><a data-view="settings" href="/settings.html">Settings</a></li>
			</ul>
			<div></div>
		</header>

		<!-- app.Main -->
		<main data-view="welcome">
			<!-- app.Views -->
			<section data-name="welcome" data-label="Welcome" data-path="/index.html" data-state="active">(...)</section>
			<section data-name="settings" data-label="Settings" data-path="/settings.html">(...)</section>
			<!-- etc -->
		</main>

		<!-- app.Main.Footer -->
		<footer>
			<div><button data-action="cancel">Cancel</button></div>
			<div></div>
			<div><button data-action="confirm">Confirm</button></div>
		</footer>

		<!-- app.Main.Dialog -->
		<dialog>
			<article>
				<button data-action="close"></button>
				<h3>Dialog Example</h3>
				<fieldset>
					<div data-name="title">
						<label data-type="title">Title</label>
						<input type="text" placeholder="Leave a Star on GitHub"/>
					</div>
					<div data-name="confirm">
						<label data-type="confirm">Confirm?</label>
						<input type="checkbox"/>
					</div>
				</fieldset>
				<footer>
					<div><button data-action="cancel">Cancel</button></div>
					<div></div>
					<div><button data-action="confirm">Confirm</button></div>
				</footer>
			</article>
		</dialog>
		<script src="wasm_exec.js"></script>
		<script src="wasm_init.js"></script>
	</body>
</html>
```

## WebView CSS

The [Gooey App Theme](/design) relies on the semantically enforced HTML5 code layout.
It is a ready-to-use default theme for Gooey Apps. Copy it to your App's `/public/design` folder,
modify it, and you're good to go.

Gooey expects reliable and consistent naming of attributes, elements, and classes. All
stylesheets must abide by these rules in order to be commitable upstream:

**Themes**:

- The [Gooey App Theme](/design) is classless and uses `data-` attributes to represent states.
- All App specific Themes have to be bundled via `go:embed` inside `/public/app/`.
- All App specific Themes can use CSS classes for better customizations.

**Components**:

- All [UI Components](/components/ui) use native input elements to be Web Accessibility compatible.
- All [UI Components](/components/ui) must have a [data-type or type](/components/types/Input.go) property.
- All [Layout Components](/components/layout) must have a [data-layout](/components/types/Layout.go) property.

**Component States**:

Component States are represented using `data-` attributes:

- `data-action` influences the event flow (see Reactive MVC Architecture).
- `data-layout` influences the layout flow (`grid`, `flex`, `flow`)
- `data-state` influences the activity/visibility (`active`)

## Implementation Guides

- [Component Implementation Guide](/docs/component-implementation-guide.md) documents how to implement components and how the Component graph works.
- [View Implementation Guide](/docs/view-implementation-guide.md) documents the View layer of the loop.
- [Controller Implementation Guide](/docs/controller-implementation-guide.md) documents the Controller layer of the loop.
- [Backend Implementation Guide](/docs/backend-implementation-guide.md) documents the shared-schema Backend that the Client talks to.

