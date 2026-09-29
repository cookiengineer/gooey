
# Component Implementation Guide

This guide explains how to implement a custom [Component](/components/Component.go) for a
Gooey App, how the Component graph is assembled, and which conventions a Component must follow
to be compatible with nesting and the graph layout.

It is one of four implementation guides; the others are:

- [view-implementation-guide.md](/docs/view-implementation-guide.md)
- [controller-implementation-guide.md](/docs/controller-implementation-guide.md)
- [backend-implementation-guide.md](/docs/backend-implementation-guide.md)

**IMPORTANT**: Before reading this guide, read the [ERRATA.md](/docs/ERRATA.md) document.
It documents the quirks of running Go in a Web Browser and is relevant even for experienced
Go developers.

## Where Components sit in the Loop

The [Reactive MVC Architecture](/docs/ARCHITECTURE.md) is built from Components. Components
are the smallest elements of the DOM and are the only layer that may touch the DOM:

```text
Controller ──action──▶ Backend
    ▲                     │
  event                update
    │                     ▼
  View ◀──render──── Model / Storage
```

- A [View](/docs/view-implementation-guide.md) **is a special Component** that additionally
  has navigation semantics (`Name`, `Label`, `Path`, `Enter`, `Leave`).
- A [Controller](/docs/controller-implementation-guide.md) is **not** a Component. It drives
  Components but must not touch the DOM directly.
- Components fire events; the Controller listens and reacts.

## The "Inheritance" Structure

Go has no inheritance. Gooey uses **composition plus interfaces** instead. Understanding the
difference between the following types is essential:

| Type | Kind | Purpose |
|:-----|:-----|:--------|
| [interfaces.Component](/components/interfaces/Component.go) | interface | The structural contract every Component satisfies. |
| [interfaces.View](/components/interfaces/View.go) | interface | Embeds all `Component` methods and adds `Name`, `Label`, `Path`, `Enter`, `Leave`. |
| [interfaces.Controller](/components/interfaces/Controller.go) | interface | Separate contract (`Name`, `Enter`, `Leave`, `Update`, `Render`). Not a Component. |
| [components.Component](/components/Component.go) | concrete struct | The shared base that holds `Element`, `Content`, and `Listeners`. |
| [app.Main](/components/app/Main.go) | concrete struct | The app orchestrator. Not an `interfaces.Component`. |
| [app.View](/components/app/View.go) | concrete struct | The default View implementation. |

The relationships are:

```text
interfaces.Component
	├── implemented by components.Component (the base)
	├── implemented by every layout/content/ui component (they hold *components.Component)
	└── extended by interfaces.View
			├── implemented by app.View
			└── implemented by every custom View

interfaces.Controller        (independent)
app.Main                     (independent, orchestrates Components)
```

The concrete implementations follow a strict composition convention: a component struct has a
field named `Component` of type `*components.Component`, plus its own state:

```go
type Fieldset struct {
	Name      string                `json:"name"`
	Label     string                `json:"label"`
	Component *components.Component `json:"component"`
	fields    []*fieldset_field
}
```

The `*components.Component` base owns the DOM `Element`, the `Content` child list, and the
`Listeners` map. The component delegates event plumbing to it and implements the remaining
interface methods itself.

**IMPORTANT**: Do not Go-embed `components.Component` anonymously. The convention is a named
`Component` field so that `component.Component.Element` stays unambiguous and every component
looks the same.

## The Component Contract

[interfaces.Component](/components/interfaces/Component.go) requires eight methods:

| Method | Purpose |
|:-------|:--------|
| `Enable() bool` | Enables interactive children. Returns `true` if something changed. |
| `Disable() bool` | Disables interactive children. Returns `true` if something changed. |
| `Mount() bool` | Reads attributes, maps children, attaches listeners. Called top-down. |
| `Unmount() bool` | Unmounts children and removes listeners. Called bottom-up. |
| `SetScheduler(Scheduler)` | Receives the [interfaces.Scheduler](/components/interfaces/Scheduler.go) propagated by `app.Main` and passes it on to nested children. |
| `Query(string) Component` | Finds a nested Component by a CSS-like selector. |
| `Render() *dom.Element` | Re-renders this component and its children, returns the element. |
| `String() string` | Serializes the component for server-side rendering. |

A View must implement these **plus** `Name`, `Label`, `Path`, `Enter`, and `Leave`.

## Lifecycle: Structural vs. Navigational

Gooey distinguishes two lifecycles:

**Structural lifecycle** (every Component):

- `Mount()` — called once after the graph is parsed. Reads `data-*` attributes, maps child
  elements to Components, and attaches DOM/event listeners.
- `Unmount()` — the mirror image. Unmounts children and removes listeners.

**Navigational lifecycle** (Views only):

- `Enter()` — the View became the active View. Sets `data-state="active"`.
- `Leave()` — another View became active. Removes `data-state`.

**IMPORTANT**: The `To*` constructors **do not** mount automatically.
`Mount()` is a top-down call tree: every Component is responsible for mounting its own nested
Components. `app.Main` calls `Mount()`/`Unmount()`; `ChangeView()` calls `Enter()`/`Leave()`.

## The Component Graph

The graph is assembled by [components.Document](/components/Document.go) together with
[app.Main](/components/app/Main.go).

1. [Document.Register](/components/Document.go) maps a tag name to a constructor, usually
   wrapped with [components.WrapComponent](/components/WrapComponent.go).
2. [Document.CreateComponent](/components/Document.go) looks up the registry for an element.
3. [Document.Mount](/components/Document.go) walks the `body` children. Registered tags are
   instantiated through their constructor; unregistered tags become a plain
   `components.Component` whose descendants are mapped recursively.
4. Each component's `Mount()` then maps its own children and calls `Mount()` on them.
5. `Render()` patches the DOM top-down using the [reconciler](/components/virtual/Reconcile.go); `Unmount()` tears down bottom-up.

### Graph Compatibility Rules

For a component to be usable inside the graph and inside other components, it **must**:

1. Provide a `ToX(*dom.Element) *X` constructor that wraps an existing element.
2. Set its `Component` field and initialize `Content` to an empty slice.
3. Register itself with `document.Register("<tag>", components.WrapComponent(ToX))`.
4. In `Mount()`, map each child element to the matching component and call `child.Mount()`.
5. Implement `Query()` recursively (and self-inclusively) over `Content`.
6. In `Render()`, re-apply attributes and reconcile the rendered children with
   `components.ReconcileComponents()` / `components.ReconcileElements()`.
7. In `SetScheduler()`, forward the scheduler to the base `Component` and to every child.
8. In `Unmount()`, unmount the children and remove any DOM listeners.
9. Keep listeners attached in `Mount()`, never in `Render()`.

A minimal graph-compatible component therefore looks like this. Known tags are mapped to their
`To*` constructors (expanded in [Child Mapping Conventions](#child-mapping-conventions)); unknown
tags are wrapped in the base component:

```go
func (component *X) Mount() bool {

	if component.Component.Element == nil {
		return false
	}

	// 1. Read attributes
	name := component.Component.Element.GetAttribute("data-name")
	if name != "" {
		component.Name = name
	}

	// 2. Map children
	elements := component.Component.Element.Children()
	children := make([]interfaces.Component, 0)

	for _, element := range elements {

		switch element.TagName {
		case "TABLE":
			children = append(children, content_components.ToTable(element))
		default:
			tmp := components.NewComponent(element)
			children = append(children, &tmp)
		}

	}

	component.Content = children

	// 3. Mount children
	for _, child := range component.Content {
		child.Mount()
	}

	return true

}
```

## Layout Components vs. Leaf Components

The nesting rules differ by component role:

- **Leaf components** (`ui.Button`, `ui.Input`, `ui.Label`, ...) have no `Content` and never
  recurse. They only update themselves.
- **Layout components** (`layout.Article`, `layout.Dialog`, `layout.Header`, `layout.Aside`,
  `layout.Footer`, `content.Fieldset`, `content.Table`) own a `Content` list and map their
  children.
- **Custom app components** are usually layout components too. A custom component can map any
  registered tag, including other custom components.
- **Views** are layout components registered under the special `section` tag.

**IMPORTANT**: "Layout Component" is a structural role in this guide. It is not limited to the
`components/layout` package; `content.Fieldset` and `content.Table` are layout components as
well.

Nested structures rely on two rules:

1. A parent owns its `Content`. It decides which child elements become Components.
2. A child never mounts itself. The parent calls `Mount()`, and the child mounts its own
   children in turn.

[Component.SetContent](/components/Component.go) is the programmatic counterpart of child
mapping: it replaces `Content` for components that are composed in Go code rather than from
existing markup.

## The Layout Property

A Layout Component can expose a `Layout` field of type
[types.Layout](/components/types/Layout.go). It controls how the component arranges itself and
its children, and it is mapped to the `data-layout` HTML attribute.

The available values are:

| Constant | Value | Meaning |
|:---------|:------|:--------|
| `types.LayoutFlow` | `flow` | Normal document flow. The default for most components. |
| `types.LayoutFlex` | `flex` | Flexbox layout. |
| `types.LayoutGrid` | `grid` | CSS grid layout. |

The recommended convention is:

- default the field to `types.LayoutFlow` in the constructor,
- read `data-layout` in `Mount()` when the attribute is present,
- sync `data-layout` in `Render()`, writing it only when the value differs from the default.

```go
type Foo struct {
	Name      string                `json:"name"`
	Layout    types.Layout          `json:"layout"`
	Component *components.Component `json:"component"`
}

func (component *Foo) Mount() bool {

	if component.Component.Element != nil {

		layout := component.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			component.Layout = types.Layout(layout)
		}

		// ... map and mount children ...

		return true

	}

	return false

}

func (component *Foo) Render() *dom.Element {

	if component.Component.Element != nil {

		if component.Layout != types.LayoutFlow {
			component.Component.Element.SetAttribute("data-layout", component.Layout.String())
		} else {
			component.Component.Element.RemoveAttribute("data-layout")
		}

		// ... reconcile children ...

	}

	return component.Component.Element

}
```

The built-in components that expose `Layout` are [layout.Header](/components/layout/Header.go),
[layout.Aside](/components/layout/Aside.go), [layout.Footer](/components/layout/Footer.go),
[layout.Dialog](/components/layout/Dialog.go), [layout.Article](/components/layout/Article.go),
and [app.View](/components/app/View.go). `content.Fieldset` and `content.Table` own children but
do not expose a `Layout` field.

**IMPORTANT**: `Header`, `Aside`, and `Footer` use `types.LayoutFlex` as their default, not
`LayoutFlow`. Check the component before assuming a default.

The [Gooey App Theme](/design) styles the states with attribute selectors such as
`header[data-layout="flex"]`, `aside[data-layout="grid"]`, and `section[data-layout="flow"]`, so
the attribute is what CSS reacts to.

## Child Mapping Conventions

Mapping is done in `Mount()` by inspecting `element.TagName`. Use the matching `To*`
constructor for known tags, and wrap unknown ones. (Gooey imports the `content` package as
`content_components` in this situation to avoid confusion with the `components` package.)

```go
mapped := make([]interfaces.Component, 0)

for _, element := range component.Component.Element.Children() {

	switch element.TagName {
	case "BUTTON":
		mapped = append(mapped, ui.ToButton(element))
	case "FIELDSET":
		mapped = append(mapped, content_components.ToFieldset(element))
	case "TABLE":
		mapped = append(mapped, content_components.ToTable(element))
	case "APP-CUSTOM-COMPONENT":
		mapped = append(mapped, app_components.ToCustomComponent(element))
	default:
		tmp := components.NewComponent(element)
		mapped = append(mapped, &tmp)
	}

}

component.Content = mapped
```

**IMPORTANT**: Map a child exactly once and call `Mount()` exactly once. Double-mapping or
double-mounting attaches duplicate listeners and renders duplicate rows.

## Query Conventions

`Query()` uses a CSS-like selector language implemented by
[components/utils](/components/utils): `SplitQuery`, `MatchesQuery`, and `JoinQuery`.

- The first selector is matched against the component's **own** element. The query is
  self-including, so `controller.View.Query("section > article > table")` starts at the View.
- For every remaining selector, recurse into each child in `Content`.
- Return `component` when the single selector matches the component itself.
- Return `nil` when nothing matches.

```go
func (component *X) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if utils.MatchesQuery(component.Component.Element, selectors[0]) == true {

			tmp_query := utils.JoinQuery(selectors[1:])

			for _, content := range component.Content {

				tmp_component := content.Query(tmp_query)

				if tmp_component != nil {
					return tmp_component
				}

			}

		}

	} else if len(selectors) == 1 {

		if utils.MatchesQuery(component.Component.Element, selectors[0]) == true {
			return component
		}

	}

	return nil

}
```

Because `Query()` returns `interfaces.Component`, callers use the generic
[components.UnwrapComponent](/components/UnwrapComponent.go) helper to recover the concrete
type:

```go
table, ok := components.UnwrapComponent[*content.Table](
	controller.View.Query("section > article > table"),
)
```

## Event Conventions

Components communicate through events, never by calling Controllers directly.

1. In `Mount()`, reserve the event with `Component.InitEvent(event)`.
2. Register a callback with
   [Component.AddEventListener](/components/Component.go) and
   [components.ToEventListener](/components/EventListener.go).
3. In the callback, or from a DOM listener, call `Component.FireEventListeners(event, attributes)`.

```go
func (component *X) Mount() bool {

	if component.Component != nil {
		component.Component.InitEvent("action")
	}

	// ... add DOM listeners that call FireEventListeners ...

	return true

}
```

The callback signature is `func(string, map[string]any)`, defined by
[EventListenerCallback](/components/EventListener.go). The first argument is the event name,
the second is the attribute map. For `click` and `change`, the base `Component` refreshes the
target's attributes before firing.

The standard events are:

| Event | Fired by |
|:------|:---------|
| `click` | `ui.Button` |
| `change-value` | `ui.Input`, `ui.Checkbox`, `ui.Number`, `ui.Range`, `ui.Select`, `ui.Textarea` |
| `change-field` | `content.Fieldset` |
| `change-view` | `layout.Header`, `layout.Aside` |
| `action` | `content.Table`, `layout.Header`, `layout.Aside`, `layout.Footer`, `layout.Dialog`, custom components |

The layout components (`layout.Header`, `layout.Aside`, `layout.Footer`, `layout.Dialog`,
`content.Table`) attach a DOM `click` listener and translate a `data-action` target into an
`action` event, so the View markup only needs `data-action` attributes.

## Render and String Conventions

`Render()` must:

- return `component.Component.Element`,
- re-apply all `data-*` attributes derived from component state,
- reconcile the rendered child elements in place instead of replacing them,
- be idempotent and safe to call repeatedly,
- never attach or detach event listeners.

Prefer [components.ReconcileComponents](/components/ReconcileComponents.go) for children that
implement `interfaces.Component` and [components.ReconcileElements](/components/ReconcileElements.go)
for raw `*dom.Element` slices. Both delegate to the [components/virtual](/components/virtual/Reconcile.go)
reconciler, which preserves the identity of unchanged nodes and keeps focus, text selection and
scroll position intact. Build list items as elements and give them a stable `data-key` so the
reconciler can match them across renders; `data-id` and `id` are also honoured as fallbacks.

```go
func (component *X) Render() *dom.Element {

	if component.Component.Element != nil {

		component.Component.Element.SetAttribute("data-name", component.Name)

		components.ReconcileComponents(component.Component.Element, component.Content)

	}

	return component.Component.Element

}
```

### Reactive Invalidation

Setters that change component state should call the wrapper's `Invalidate()` method. It marks the
base `Component` dirty, bumps its revision and asks the scheduler to render it as soon as possible.

Components do not own a scheduler by themselves. [app.Main.Mount](/components/app/Main.go)
propagates its [app.Scheduler](/components/app/Scheduler.go) through the whole graph by calling
`SetScheduler()` on every Component **after** the graph has been mounted; only then can
`Invalidate()` reach the [interfaces.Scheduler](/components/interfaces/Scheduler.go). The
scheduler coalesces all invalidations in one animation frame into a single `Render()`. A Component
that is not part of an `app.Main` graph simply has no scheduler and is rendered explicitly instead.

Because the base `Component` has no reference back to its wrapper, wrapper components override
`Invalidate()` and delegate to `Component.InvalidateAs(owner)` so that the wrapper's own `Render()`
is scheduled:

```go
func (table *Table) Invalidate() {
	table.Component.InvalidateAs(table)
}

func (table *Table) SetDataset(dataset data.Dataset) {
	table.Dataset = &dataset
	table.Invalidate()
}
```

The base [components.Component](/components/Component.go) also provides the building blocks for
custom scheduling logic: `MarkDirty()` flags a change without scheduling, `ClearDirty()` resets the
flag after a render, `IsDirty()` and `Revision()` report the pending state, `SetScheduler()` stores
and propagates a scheduler, and `Schedule(owner)` enqueues a specific owner.

`String()` serializes the component and its `Content` into HTML. It is used for server-side
rendering and should produce markup that is semantically equivalent to `Render()`.

## Enable and Disable Conventions

`Enable()`/`Disable()` propagate to interactive children and return `true` when something
changed. Leaf components toggle the `disabled` attribute; layout components iterate their
`Content`.

```go
func (component *X) Disable() bool {

	var result bool

	for _, content := range component.Content {
		if content.Disable() == true {
			result = true
		}
	}

	return result

}
```

## Constructors and Registration

Two constructor conventions exist across Gooey:

- `NewX(...)` builds a detached element from code. In the `layout`, `content`, `ui`, and
  custom component packages it returns a **value**; `app.View` and `app.Main` return pointers.
- `ToX(element *dom.Element) *X` always wraps an existing element and returns a **pointer**.

Both initialize the `Component` field. Register the `To*` constructor in a `RegisterTo`
function:

```go
package components

import "github.com/cookiengineer/gooey/components"

func RegisterTo(document *components.Document) {

	document.Register("app-custom-component", components.WrapComponent(ToCustomComponent))

}
```

Then register the package from `app/main.go` before `main.Mount()`:

```go
app_components.RegisterTo(main.Document)
```

**IMPORTANT**: Registration order relative to `Mount()` matters. Register everything before
`app.Main.Mount()` so the graph parser can find the constructors.

## Where Custom Components Live

Following the recommended App layout (see the
[View guide](/docs/view-implementation-guide.md#recommended-app-file-layout)):

```text
app/
├── components/
│   ├── RegisterTo.go
│   └── CustomComponent.go
└── public/
	└── app/
		└── components/
			└── CustomComponent.css
```

Custom components must have the `//go:build wasm` build tag because they use DOM bindings.

## Full Example

See
[examples/components/app-components/components/CustomComponent.go](/examples/components/app-components/components/CustomComponent.go)
for a complete layout component, and
[examples/components/app-components/components/RegisterTo.go](/examples/components/app-components/components/RegisterTo.go)
for its registration. The [app-views example](/examples/components/app-views/views/Settings.go)
shows how a custom View maps an app-specific component tag in `Mount()`.

## Available Components

**components** (`github.com/cookiengineer/gooey/components`):

- [Component](/components/Component.go)
- [Document](/components/Document.go)
- [EventListener](/components/EventListener.go)
- Reconciliation helpers: [ReconcileComponents](/components/ReconcileComponents.go),
  [ReconcileElements](/components/ReconcileElements.go)
- Helpers: [ComponentConstructor](/components/ComponentConstructor.go),
  [UnwrapComponent](/components/UnwrapComponent.go),
  [WrapComponent](/components/WrapComponent.go)

**components/app**:

- [app.Client](/components/app/Client.go)
- [app.ClientListener](/components/app/ClientListener.go)
- [app.Controller](/components/app/Controller.go)
- [app.Main](/components/app/Main.go)
- [app.Scheduler](/components/app/Scheduler.go) coalesces invalidated Components once per animation frame
- [app.Storage](/components/app/Storage.go) with the reactive `Get()`, `Revision()`, `Update()` and `Subscribe()` methods
- [app.View](/components/app/View.go)
- Controller helpers: [ControllerConstructor](/components/app/ControllerConstructor.go),
  [UnwrapController](/components/app/UnwrapController.go),
  [WrapController](/components/app/WrapController.go)
- View helpers: [ViewConstructor](/components/app/ViewConstructor.go),
  [UnwrapView](/components/app/UnwrapView.go), [WrapView](/components/app/WrapView.go)

**components/interfaces**:

- [interfaces.Component](/components/interfaces/Component.go)
- [interfaces.Controller](/components/interfaces/Controller.go)
- [interfaces.View](/components/interfaces/View.go)
- [interfaces.Scheduler](/components/interfaces/Scheduler.go)

**components/content**:

- [content.Fieldset](/components/content/Fieldset.go) fires a `change-field` event
- [content.LineChart](/components/content/LineChart.go)
- [content.PieChart](/components/content/PieChart.go)
- [content.Table](/components/content/Table.go) fires an `action` event; rows carry a stable
  `data-key` derived from the `Identifier` field (the `data-identifier` attribute, default `"id"`)
  so the reconciler can match them, and `SelectedKeys()` returns the keys of the selected rows

**components/data** (helpers used by `content.Table`):

- [data.Data](/components/data/Data.go)
- [data.Dataset](/components/data/Dataset.go)

**components/reactive** (pure-Go core, unit-testable without WebASM):

- [reactive.Store](/components/reactive/Store.go) observable key/value store used by `app.Storage`
- [reactive.Queue](/components/reactive/Queue.go) coalescing FIFO queue used by `app.Scheduler`

**components/virtual** (DOM reconciler):

- [virtual.Reconcile](/components/virtual/Reconcile.go) patches children in place
- [virtual.Patch](/components/virtual/Patch.go), [virtual.Move](/components/virtual/Move.go)
- [virtual.Element](/components/virtual/Element.go) adapts `*dom.Element` for reconciliation

**components/ui**:

- [ui.Button](/components/ui/Button.go) fires a `click` event
- [ui.Checkbox](/components/ui/Checkbox.go) fires a `change-value` event
- [ui.Input](/components/ui/Input.go) fires a `change-value` event
- [ui.Label](/components/ui/Label.go)
- [ui.Number](/components/ui/Number.go) fires a `change-value` event
- [ ] ui.Radio fires a `change-value` event (not implemented yet)
- [ui.Range](/components/ui/Range.go) fires a `change-value` event
- [ui.Select](/components/ui/Select.go) fires a `change-value` event
- [ui.Textarea](/components/ui/Textarea.go) fires a `change-value` event

**components/layout**:

- [layout.Article](/components/layout/Article.go)
- [layout.Aside](/components/layout/Aside.go) fires an `action` and a `change-view` event
- [layout.Dialog](/components/layout/Dialog.go) fires an `action` event
- [layout.Footer](/components/layout/Footer.go) fires an `action` event
- [layout.Header](/components/layout/Header.go) fires an `action` and a `change-view` event

## Component Checklist

- [ ] File has the `//go:build wasm` build tag.
- [ ] Struct has a `Component *components.Component` field.
- [ ] `ToX(...)` returns a pointer and initializes `Component` and `Content`.
- [ ] Layout components expose `Layout types.Layout` and sync `data-layout` in `Render()`.
- [ ] `Mount()` reads attributes, maps children, and calls `child.Mount()`.
- [ ] `Unmount()` unmounts children and removes DOM listeners.
- [ ] `SetScheduler()` forwards the scheduler to the base `Component` and to every child.
- [ ] `Query()` is self-including and recurses into `Content`.
- [ ] `Render()` returns `Component.Element` and reconciles children.
- [ ] `String()` produces equivalent markup.
- [ ] Events are reserved with `InitEvent()` and fired via `FireEventListeners()`.
- [ ] Listeners are attached in `Mount()`, never in `Render()`.
- [ ] The component is registered before `app.Main.Mount()`.
- [ ] App-owned CSS lives in `public/app/components/<Name>.css`.

## References

- [interfaces.Component](/components/interfaces/Component.go)
- [interfaces.View](/components/interfaces/View.go)
- [components.Component](/components/Component.go)
- [components.Document](/components/Document.go)
- [components.EventListener](/components/EventListener.go)
- [components.utils](/components/utils)
- [app.Main](/components/app/Main.go)
- [app-components example](/examples/components/app-components)
- [app-views example](/examples/components/app-views)
- [ARCHITECTURE.md](/docs/ARCHITECTURE.md)
