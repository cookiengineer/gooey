
# Components

- [Component](/components/Component.go)
- [Document](/components/Document.go)
- [EventListener](/components/EventListener.go)

**Component Helpers**:

- [ComponentConstructor](/components/ComponentConstructor.go)
- [UnwrapComponent](/components/UnwrapComponent.go)
- [WrapComponent](/components/WrapComponent.go)

## components/app

- [app.Client](/components/app/Client.go)
- [app.ClientListener](/components/app/ClientListener.go)
- [app.Controller](/components/app/Controller.go)
- [app.Main](/components/app/Main.go)
- [app.Storage](/components/app/Storage.go)
- [app.View](/components/app/View.go)

**Controller Helpers**:

- [ControllerConstructor](/components/app/ControllerConstructor.go)
- [UnwrapController](/components/app/UnwrapController.go)
- [WrapController](/components/app/WrapController.go)

**View Helpers**:

- [ViewConstructor](/components/app/ViewConstructor.go)
- [UnwrapView](/components/app/UnwrapView.go)
- [WrapView](/components/app/WrapView.go)

## components/interfaces

- [interfaces.Component](/components/interfaces/Component.go)
- [interfaces.Controller](/components/interfaces/Controller.go)
- [interfaces.View](/components/interfaces/View.go)

## components/content

- [content.Fieldset](/components/content/Fieldset.go) fires a `change-field` event
- [content.LineChart](/components/content/LineChart.go)
- [content.PieChart](/components/content/PieChart.go)
- [content.Table](/components/content/Table.go)

## components/ui

- [ui.Button](/components/ui/Button.go) fires a `click` event
- [ui.Checkbox](/components/ui/Checkbox.go) fires a `change-value` event
- [ui.Input](/components/ui/Input.go) fires a `change-value` event
- [ui.Label](/components/ui/Label.go)
- [ui.Number](/components/ui/Number.go) fires a `change-value` event
- [ ] [ui.Radio](/components/ui/Radio.go) fires a `change-value` event
- [ui.Range](/components/ui/Range.go) fires a `change-value` event
- [ui.Select](/components/ui/Select.go) fires a `change-value` event
- [ui.Textarea](/components/ui/Textarea.go) fires a `change-value` event

## components/layout

- [ ] [layout.Aside](/components/layout/Aside.go)
- [layout.Dialog](/components/layout/Dialog.go)
- [layout.Footer](/components/layout/Footer.go) fires an `action` event
- [layout.Header](/components/layout/Header.go) fires an `action` and a `change-view` event
