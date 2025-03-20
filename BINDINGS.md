
# Bindings

- [x] [Document](/pkg/Document.go)
- [x] [Screen](/pkg/Screen.go)
- [x] [ScreenOrientation](/pkg/ScreenOrientation.go)
- [x] [Window](/pkg/Window.go)

**animations**

- [x] [animations/CancelAnimationFrame](/pkg/animations/CancelAnimationFrame.go)
- [x] [animations/RequestAnimationFrame](/pkg/animations/RequestAnimationFrame.go)

**console**

- [x] [console/Clear](/pkg/console/Clear.go)
- [x] [console/Error](/pkg/console/Error.go)
- [x] [console/Group](/pkg/console/Group.go)
- [x] [console/GroupEnd](/pkg/console/GroupEnd.go)
- [x] [console/Info](/pkg/console/Info.go)
- [x] [console/Log](/pkg/console/Log.go)
- [x] [console/Warn](/pkg/console/Warn.go)

**crypto**

The Web Crypto API has been split up into separate sub-packages, due to the `algorithm` parameters
object not being implementable as a reusable struct that would make sense.

- [x] [crypto/GetRandomValues](/pkg/crypto/GetRandomValues.go)
- [x] [crypto/RandomUUID](/pkg/crypto/RandomUUID.go)

**crypto/aescbc**

- [x] [crypto/aescbc/CryptoKey](/pkg/crypto/aescbc/CryptoKey.go)
- [x] [crypto/aescbc/CryptoKeyType](/pkg/crypto/aescbc/CryptoKeyType.go)
- [x] [crypto/aescbc/Decrypt](/pkg/crypto/aescbc/Decrypt.go)
- [x] [crypto/aescbc/Encrypt](/pkg/crypto/aescbc/Encrypt.go)
- [x] [crypto/aescbc/ExportKey](/pkg/crypto/aescbc/ExportKey.go)
- [x] [crypto/aescbc/GenerateKey](/pkg/crypto/aescbc/GenerateKey.go)
- [x] [crypto/aescbc/ImportKey](/pkg/crypto/aescbc/ImportKey.go)

**crypto/aesctr**

- [x] [crypto/aesctr/CryptoKey](/pkg/crypto/aesctr/CryptoKey.go)
- [x] [crypto/aesctr/CryptoKeyType](/pkg/crypto/aesctr/CryptoKeyType.go)
- [x] [crypto/aesctr/Decrypt](/pkg/crypto/aesctr/Decrypt.go)
- [x] [crypto/aesctr/Encrypt](/pkg/crypto/aesctr/Encrypt.go)
- [x] [crypto/aesctr/ExportKey](/pkg/crypto/aesctr/ExportKey.go)
- [x] [crypto/aesctr/GenerateKey](/pkg/crypto/aesctr/GenerateKey.go)
- [x] [crypto/aesctr/ImportKey](/pkg/crypto/aesctr/ImportKey.go)

**crypto/aesgcm**

- [x] [crypto/aesgcm/CryptoKey](/pkg/crypto/aesgcm/CryptoKey.go)
- [x] [crypto/aesgcm/CryptoKeyType](/pkg/crypto/aesgcm/CryptoKeyType.go)
- [x] [crypto/aesgcm/Decrypt](/pkg/crypto/aesgcm/Decrypt.go)
- [x] [crypto/aesgcm/Encrypt](/pkg/crypto/aesgcm/Encrypt.go)
- [x] [crypto/aesgcm/ExportKey](/pkg/crypto/aesgcm/ExportKey.go)
- [x] [crypto/aesgcm/GenerateKey](/pkg/crypto/aesgcm/GenerateKey.go)
- [x] [crypto/aesgcm/ImportKey](/pkg/crypto/aesgcm/ImportKey.go)

**dom**

- [x] [dom/Element](/pkg/dom/Element.go)
- [x] [dom/Event](/pkg/dom/Event.go)
- [x] [dom/EventListener](/pkg/dom/EventListener.go)
- [x] [dom/EventPhase](/pkg/dom/EventPhase.go)
- [x] [dom/EventType](/pkg/dom/EventType.go)

**fetch**

Note: If you run into problems, use the [Synchronous XMLHttpRequest](/pkg/xhr/XMLHttpRequest_sync.go) APIs instead.

- [x] [fetch/Fetch](/pkg/fetch/Fetch.go) [2]
- [x] [fetch/Headers](/pkg/fetch/Headers.go)
- [x] [fetch/Request](/pkg/fetch/Request.go) (or `RequestInit` object)
- [x] [fetch/Response](/pkg/fetch/Response.go)

Fetch RequestInit Properties:

- [x] [fetch/Cache](/pkg/fetch/Cache.go)
- [x] [fetch/Credentials](/pkg/fetch/Credentials.go)
- [x] [fetch/Method](/pkg/fetch/Method.go)
- [x] [fetch/Mode](/pkg/fetch/Mode.go)
- [x] [fetch/Redirect](/pkg/fetch/Redirect.go)
- [x] `Referrer` has to be a `string` due to arbitrary URL values.
- [x] [fetch/ReferrerPolicy](/pkg/fetch/ReferrerPolicy.go)

**history**

- [x] [history/EventListener](/pkg/history/EventListener.go)
- [x] [history/EventPhase](/pkg/history/EventPhase.go)
- [x] [history/EventType](/pkg/history/EventType.go)
- [x] [history/History](/pkg/history/History.go)
- [x] [history/HistoryState](/pkg/history/HistoryState.go)
- [x] [history/PopStateEvent](/pkg/history/PopStateEvent.go)

**location**

- [x] [location/Location](/pkg/location/Location.go)

**navigator**

The Navigator API is split up into separate sub-packages, due to most of the features not
being available in at least one web browser.

- [x] [navigator/Navigator](/pkg/navigator/Navigator.go)

**navigator/geolocation**

- [x] [navigator/geolocation/Geolocation](/pkg/navigator/geolocation/Geolocation.go)
- [x] [navigator/geolocation/GeolocationPosition](/pkg/navigator/geolocation/GeolocationPosition.go)
- [x] [navigator/geolocation/GeolocationPositionError](/pkg/navigator/geolocation/GeolocationPositionError.go)
- [x] [navigator/geolocation/GeolocationPositionOptions](/pkg/navigator/geolocation/GeolocationPositionOptions.go) [1]

**storages**

- [x] [storages/LocalStorage](/pkg/storages/LocalStorage.go)
- [x] [storages/SessionStorage](/pkg/storages/SessionStorage.go)

**timers**

- [x] [timers/ClearInterval](/pkg/timers/ClearInterval.go)
- [x] [timers/ClearTimeout](/pkg/timers/ClearTimeout.go)
- [x] [timers/SetInterval](/pkg/timers/SetInterval.go)
- [x] [timers/SetTimeout](/pkg/timers/SetTimeout.go)

**xhr**

- [x] [xhr/Method](/pkg/xhr/Method.go)
- [x] [xhr/XMLHttpRequest](/pkg/xhr/XMLHttpRequest.go) [2]
- [x] Synchronous [xhr/XMLHttpRequest](/pkg/xhr/XMLHttpRequest_sync.go)

--------

[1] This feature is implemented, but not supported across all Browsers. It is disabled to prevent WebASM runtime errors that are irrecoverable.

[2] This feature is implemented asynchronously and uses a go `chan`. It only works with `tinygo` as a compiler as of now. If your WebASM binary
    hangs when using this, use the synchronous XMLHttpRequest APIs instead.

