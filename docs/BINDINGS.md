
# Bindings

- [x] [Document](/source/Document.go)
- [x] [Screen](/source/Screen.go)
- [x] [ScreenOrientation](/source/ScreenOrientation.go)
- [x] [Window](/source/Window.go)

**animations**

- [x] [animations/CancelAnimationFrame](/source/animations/CancelAnimationFrame.go)
- [x] [animations/RequestAnimationFrame](/source/animations/RequestAnimationFrame.go)

**console**

- [x] [console/Clear](/source/console/Clear.go)
- [x] [console/Error](/source/console/Error.go)
- [x] [console/Group](/source/console/Group.go)
- [x] [console/GroupEnd](/source/console/GroupEnd.go)
- [x] [console/Info](/source/console/Info.go)
- [x] [console/Log](/source/console/Log.go)
- [x] [console/Warn](/source/console/Warn.go)

**crypto**

The Web Crypto API has been split up into separate sub-packages, due to the `algorithm` parameters
object not being implementable as a reusable struct that would make sense.

- [x] [crypto/GetRandomValues](/source/crypto/GetRandomValues.go)
- [x] [crypto/RandomUUID](/source/crypto/RandomUUID.go)

- [x] [crypto/aescbc/CryptoKey](/source/crypto/aescbc/CryptoKey.go)
- [x] [crypto/aescbc/Encrypt](/source/crypto/aescbc/Encrypt.go)
- [x] [crypto/aescbc/Decrypt](/source/crypto/aescbc/Decrypt.go)
- [x] [crypto/aescbc/GenerateKey](/source/crypto/aescbc/GenerateKey.go)

- [ ] [crypto/aesgcm/CryptoKey](/source/crypto/aescbc/CryptoKey.go)
- [ ] [crypto/aesgcm/Encrypt](/source/crypto/aescbc/Encrypt.go)
- [ ] [crypto/aesgcm/Decrypt](/source/crypto/aescbc/Decrypt.go)
- [ ] [crypto/aesgcm/Params](/source/crypto/aescbc/Params.go)

**dom**

- [x] [dom/Element](/source/dom/Element.go)
- [x] [dom/Event](/source/dom/Event.go)
- [x] [dom/EventListener](/source/dom/EventListener.go)
- [x] [dom/EventPhase](/source/dom/EventPhase.go)
- [x] [dom/EventType](/source/dom/EventType.go)

**fetch**

Note: If you run into problems, use the [Synchronous XMLHttpRequest](/source/xhr/XMLHttpRequest_sync.go) APIs instead.

- [x] [fetch/Fetch](/source/fetch/Fetch.go) [2]
- [x] [fetch/Headers](/source/fetch/Headers.go)
- [x] [fetch/Request](/source/fetch/Request.go) (or `RequestInit` object)
- [x] [fetch/Response](/source/fetch/Response.go)

Fetch RequestInit Properties:

- [x] [fetch/Cache](/source/fetch/Cache.go)
- [x] [fetch/Credentials](/source/fetch/Credentials.go)
- [x] [fetch/Method](/source/fetch/Method.go)
- [x] [fetch/Mode](/source/fetch/Mode.go)
- [x] [fetch/Redirect](/source/fetch/Redirect.go)
- [x] `Referrer` has to be a `string` due to arbitrary URL values.
- [x] [fetch/ReferrerPolicy](/source/fetch/ReferrerPolicy.go)

**history**

- [x] [history/EventListener](/source/history/EventListener.go)
- [x] [history/EventPhase](/source/history/EventPhase.go)
- [x] [history/EventType](/source/history/EventType.go)
- [x] [history/History](/source/history/History.go)
- [x] [history/HistoryState](/source/history/HistoryState.go)
- [x] [history/PopStateEvent](/source/history/PopStateEvent.go)

**location**

- [x] [location/Location](/source/location/Location.go)

**navigator**

The Navigator API is split up into separate sub-packages, due to most of the features not
being available in at least one web browser.

- [x] [navigator/Navigator](/source/navigator/Navigator.go)
- [x] [navigator/geolocation/Geolocation](/source/navigator/geolocation/Geolocation.go)
- [x] [navigator/geolocation/GeolocationPosition](/source/navigator/geolocation/GeolocationPosition.go)
- [x] [navigator/geolocation/GeolocationPositionError](/source/navigator/geolocation/GeolocationPositionError.go)
- [x] [navigator/geolocation/GeolocationPositionOptions](/source/navigator/geolocation/GeolocationPositionOptions.go) [1]

**storages**

- [x] [storages/LocalStorage](/source/storages/LocalStorage.go)
- [x] [storages/SessionStorage](/source/storages/SessionStorage.go)

**timers**

- [x] [timers/ClearInterval](/source/timers/ClearInterval.go)
- [x] [timers/ClearTimeout](/source/timers/ClearTimeout.go)
- [x] [timers/SetInterval](/source/timers/SetInterval.go)
- [x] [timers/SetTimeout](/source/timers/SetTimeout.go)

**xhr**

- [x] [xhr/Method](/source/xhr/Method.go)
- [x] [xhr/XMLHttpRequest](/source/xhr/XMLHttpRequest.go) [2]
- [x] Synchronous [xhr/XMLHttpRequest](/source/xhr/XMLHttpRequest_sync.go)

--------

[1] This feature is implemented, but not supported across all Browsers. It is disabled to prevent WebASM runtime errors that are irrecoverable.

[2] This feature is implemented asynchronously and uses a go `chan`. It only works with `tinygo` as a compiler as of now. If your WebASM binary
    hangs when using this, use the synchronous XMLHttpRequest APIs instead.

