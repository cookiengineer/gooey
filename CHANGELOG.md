
## [v0.0.6] - UNRELEASED

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
