
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

- Added [ui.Button](/components/ui/Button.go) methods `SetAction(action string)`, `SetLabel(label string)`
- Added [ui.Checkbox](/components/ui/Checkbox.go) methods `SetLabel(label string)`, `SetValue(value bool)`
- Added [ui.Input](/components/ui/Input.go) methods `SetLabel(label string)`, `SetValue(value string)`
- Added [ui.Label](/components/ui/Label.go) methods `SetLabel(label string)`, `SetType(type string)`
- Added [ui.Number](/components/ui/Number.go) methods `SetLabel(label string)`, `SetMinMaxStep(min_value int, max_value int, step int)`, `SetValue(value int)`
- Added [ui.Range](/components/ui/Range.go) methods `SetLabel(label string)`, `SetMinMaxStep(min_value int, max_value int, step int)`, `SetValue(value int)`
- Added [ui.Select](/components/ui/Select.go) methods `SetLabel(label string)`, `SetValue(value string)`, `SetValues(values []string)`
- Added [ui.Textarea](/components/ui/Textarea.go) methods `SetLabel(label string)`, `SetValue(value string)`

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
