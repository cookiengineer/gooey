
## [v0.0.8] - UNRELEASED

### Changed

- Changed [console.Console](/bindings/console) to method `GetConsole() *Console` for better API docs
- Changed [bindings.Window](/bindings/Window.go) to method `GetWindow() *Window` for better API docs
- Changed [dom.Document](/bindings/dom/Document.go) to method `GetDocument() *Document` for better API docs
- Changed [history.History](/bindings/history/History.go) to method `GetHistory() *History` for better API docs
- Changed [location.Location](/bindings/location/Location.go) to method `GetLocation() *Location` for better API docs
- Changed [navigator.Navigator](/bindings/navigator/Navigator.go) to method `GetNavigator() *Navigator` for better API docs
- Changed [storages.LocalStorage](/bindings/storages/LocalStorage.go) to method `GetLocalStorage() *Storage` for better API docs
- Changed [storages.SessionStorage](/bindings/storages/SessionStorage.go) to method `GetSessionStorage() *Storage` for better API docs
- Changed [app.Main](/components/app/Main.go) method `ChangeView()` which now calls both the `View`'s and the `Controller`'s `Leave()` and `Enter()` methods correctly
- Changed [app.View](/components/app/View.go) method `Mount()` which now uses the `app.Main.Document`'s Component Registry
- Added [interfaces.Controller](/components/interfaces/Controller.go) method `Enter() bool`
- Added [interfaces.Controller](/components/interfaces/Controller.go) method `Leave() bool`

### Added

- Added [components.Document](/components/Document.go) method `CreateComponent(*dom.Element) interfaces.Component`
- Added [components.Document](/components/Document.go) method `IsRegistered(tagname string) bool`
- Added [app.Controller](/components/app/Controller.go) method `Enter() bool`
- Added [app.Controller](/components/app/Controller.go) method `Leave() bool`

## [v0.0.7] - 2025-11-03

### Changed

- Fixes for [app.View](/design/components/app/View.css) stylesheets
- Fixes for [content.Fieldset](/design/components/content/Fieldset.css) stylesheets
- Fixes for [layout.Header](/design/components/layout/Header.css) stylesheets
- Fixes for [layout.Aside](/design/components/layout/Aside.css) stylesheets
- Fixes for [layout.Dialog](/design/components/layout/Dialog.css) stylesheets
- Fixes for [layout.Footer](/design/components/layout/Footer.css) stylesheets

### Added

- Added [app.Main](/components/app/Main.go) property `Aside *layout.Aside`
- Added [layout.Aside](/components/layout/Aside.go) component as an alternative to `layout.Header`

### Removed

- Removed `layout.Header` property `View string`

## [v0.0.6] - 2025-10-15

### Changed

The `Component.Mount()` method is now called separately *after* the Component Graph has been parsed
by the [components.Document](/components/Document.go) and is now a top-down call tree. Each Component
is responsible for Mounting and Unmounting its own nested Components.

- Changed `gooey/components/app.To*` behavior (no auto-mount)
- Changed `gooey/components/content.To*` behavior (no auto-mount)
- Changed `gooey/components/layout.To*` behavior (no auto-mount)
- Changed `gooey/components/ui.To*` behavior (no auto-mount)

### Added

- Added [layout.Article](/components/layout/Article.go) method `SetContent([]interfaces.Component)`
- Added [layout.Dialog](/components/layout/Dialog.go) method `SetContent([]interfaces.Component)`
- Added [layout.Dialog](/components/layout/Dialog.go) method `SetFooter(*layout.Footer)`
- Added [layout.Dialog](/components/layout/Dialog.go) method `SetTitle(string)`
- Added [layout.Footer](/components/layout/Footer.go) method `SetContentCenter([]interfaces.Component)`
- Added [layout.Footer](/components/layout/Footer.go) method `SetContentLeft([]interfaces.Component)`
- Added [layout.Footer](/components/layout/Footer.go) method `SetContentRight([]interfaces.Component)`
- Added [layout.Header](/components/layout/Header.go) method `SetContentLeft([]interfaces.Component)`
- Added [layout.Header](/components/layout/Header.go) method `SetContentRight([]interfaces.Component)`
- Added [ui.Button](/components/ui/Button.go) method `SetAction(string)`
- Added [ui.Button](/components/ui/Button.go) method `SetLabel(string)`
- Added [ui.Checkbox](/components/ui/Checkbox.go) method `SetLabel(string)`
- Added [ui.Checkbox](/components/ui/Checkbox.go) method `SetValue(bool)`
- Added [ui.Input](/components/ui/Input.go) method `SetLabel(string)`
- Added [ui.Input](/components/ui/Input.go) method `SetValue(string)`
- Added [ui.Label](/components/ui/Label.go) method `SetLabel(string)`
- Added [ui.Label](/components/ui/Label.go) method `SetType(string)`
- Added [ui.Number](/components/ui/Number.go) method `SetLabel(string)`
- Added [ui.Number](/components/ui/Number.go) method `SetMinMaxStep(int, int, int)`
- Added [ui.Number](/components/ui/Number.go) method `SetValue(int)`
- Added [ui.Range](/components/ui/Range.go) method `SetLabel(string)`
- Added [ui.Range](/components/ui/Range.go) method `SetMinMaxStep(int, int, int)`
- Added [ui.Range](/components/ui/Range.go) method `SetValue(int)`
- Added [ui.Select](/components/ui/Select.go) method `SetLabel(string)`
- Added [ui.Select](/components/ui/Select.go) method `SetValue(string)`
- Added [ui.Select](/components/ui/Select.go) method `SetValues([]string)`
- Added [ui.Textarea](/components/ui/Textarea.go) method `SetLabel(string)`
- Added [ui.Textarea](/components/ui/Textarea.go) method `SetValue(string)`

## [v0.0.5] - 2025-10-10

### Changed

- Moved `gooey/interfaces` to `gooey/components/interfaces`
- Moved `gooey/types` to `gooey/components/types`
- Changed [interfaces.Controller](/components/interfaces/Controller.go)
- Changed [interfaces.View](/components/interfaces/View.go) to fulfill also [interfaces.Component](/components/interfaces/Component.go)
- Renamed `components.Wrap` to [components.WrapComponent](/components/WrapComponent.go) helper
- Renamed `components.Unwrap` to [components.UnwrapComponent](/components/app/UnwrapComponent.go) helper

### Added

- Added [app.WrapController](/components/app/WrapController.go) helper
- Added [app.UnwrapController](/components/app/UnwrapController.go) helper
- Added [app.WrapView](/components/app/WrapView.go) helper
- Added [app.UnwrapView](/components/app/UnwrapView.go) helper

## [v0.0.4] - 2025-10-02

Initial release
