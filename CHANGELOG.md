
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
