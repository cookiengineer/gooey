
# Bindings

- [bindings.Document](/bindings/Document.go)
- [bindings.Screen](/bindings/Screen.go)
- [bindings.ScreenOrientation](/bindings/ScreenOrientation.go)
- [bindings.Window](/bindings/Window.go)

**animations**

- [animations.CancelAnimationFrame](/bindings/animations/CancelAnimationFrame.go)
- [animations.RequestAnimationFrame](/bindings/animations/RequestAnimationFrame.go)

**console**

- [console.Clear](/bindings/console/Clear.go)
- [console.Error](/bindings/console/Error.go)
- [console.Group](/bindings/console/Group.go)
- [console.GroupEnd](/bindings/console/GroupEnd.go)
- [console.Info](/bindings/console/Info.go)
- [console.Log](/bindings/console/Log.go)
- [console.Warn](/bindings/console/Warn.go)

**cookiestore**

The Cookiestore API is currently only fully supported in Chromium-based Web Views.

- [cookiestore.Cookie](/bindings/cookiestore/Cookie.go)
- [cookiestore.Delete](/bindings/cookiestore/Delete.go)
- [cookiestore.DeleteOptions](/bindings/cookiestore/DeleteOptions.go)
- [cookiestore.Get](/bindings/cookiestore/Get.go)
- [cookiestore.GetOptions](/bindings/cookiestore/GetOptions.go)
- [cookiestore.SameSite](/bindings/cookiestore/SameSite.go)
- [cookiestore.Set](/bindings/cookiestore/Set.go)
- [cookiestore.SetOptions](/bindings/cookiestore/SetOptions.go)

**crypto**

The Web Crypto API has been split up into separate sub-packages, due to the `algorithm` parameters
object not being implementable as a reusable struct that would make sense.

- [crypto.GetRandomValues](/bindings/crypto/GetRandomValues.go)
- [crypto.RandomUUID](/bindings/crypto/RandomUUID.go)

**crypto/aescbc**

- [aescbc.CryptoKey](/bindings/crypto/aescbc/CryptoKey.go)
- [aescbc.CryptoKeyType](/bindings/crypto/aescbc/CryptoKeyType.go)
- [aescbc.Decrypt](/bindings/crypto/aescbc/Decrypt.go)
- [aescbc.Encrypt](/bindings/crypto/aescbc/Encrypt.go)
- [aescbc.ExportKey](/bindings/crypto/aescbc/ExportKey.go)
- [aescbc.GenerateKey](/bindings/crypto/aescbc/GenerateKey.go)
- [aescbc.ImportKey](/bindings/crypto/aescbc/ImportKey.go)

**crypto/aesctr**

- [aesctr.CryptoKey](/bindings/crypto/aesctr/CryptoKey.go)
- [aesctr.CryptoKeyType](/bindings/crypto/aesctr/CryptoKeyType.go)
- [aesctr.Decrypt](/bindings/crypto/aesctr/Decrypt.go)
- [aesctr.Encrypt](/bindings/crypto/aesctr/Encrypt.go)
- [aesctr.ExportKey](/bindings/crypto/aesctr/ExportKey.go)
- [aesctr.GenerateKey](/bindings/crypto/aesctr/GenerateKey.go)
- [aesctr.ImportKey](/bindings/crypto/aesctr/ImportKey.go)

**crypto/aesgcm**

- [aesgcm.CryptoKey](/bindings/crypto/aesgcm/CryptoKey.go)
- [aesgcm.CryptoKeyType](/bindings/crypto/aesgcm/CryptoKeyType.go)
- [aesgcm.Decrypt](/bindings/crypto/aesgcm/Decrypt.go)
- [aesgcm.Encrypt](/bindings/crypto/aesgcm/Encrypt.go)
- [aesgcm.ExportKey](/bindings/crypto/aesgcm/ExportKey.go)
- [aesgcm.GenerateKey](/bindings/crypto/aesgcm/GenerateKey.go)
- [aesgcm.ImportKey](/bindings/crypto/aesgcm/ImportKey.go)

**dom**

- [dom.Element](/bindings/dom/Element.go)
- [dom.Event](/bindings/dom/Event.go)
- [dom.EventListener](/bindings/dom/EventListener.go)
- [dom.EventPhase](/bindings/dom/EventPhase.go)
- [dom.EventType](/bindings/dom/EventType.go)

**fetch**

Note: If you run into problems, use the [Synchronous XMLHttpRequest](/bindings/xhr/XMLHttpRequest_sync.go) APIs instead.

- [fetch.Fetch](/bindings/fetch/Fetch.go) [2]
- [fetch.Headers](/bindings/fetch/Headers.go)
- [fetch.Request](/bindings/fetch/Request.go) (or `RequestInit` object)
- [fetch.Response](/bindings/fetch/Response.go)

Fetch RequestInit Properties:

- [fetch.Cache](/bindings/fetch/Cache.go)
- [fetch.Credentials](/bindings/fetch/Credentials.go)
- [fetch.Method](/bindings/fetch/Method.go)
- [fetch.Mode](/bindings/fetch/Mode.go)
- [fetch.Redirect](/bindings/fetch/Redirect.go)
- `Referrer` has to be a `string` due to arbitrary URL values.
- [fetch.ReferrerPolicy](/bindings/fetch/ReferrerPolicy.go)

**history**

- [history.EventListener](/bindings/history/EventListener.go)
- [history.EventPhase](/bindings/history/EventPhase.go)
- [history.EventType](/bindings/history/EventType.go)
- [history.History](/bindings/history/History.go)
- [history.HistoryState](/bindings/history/HistoryState.go)
- [history.PopStateEvent](/bindings/history/PopStateEvent.go)

**location**

- [location.Location](/bindings/location/Location.go)

**navigator**

The Navigator API is split up into separate sub-packages, due to most of the features not
being available in at least one web browser.

- [navigator.Navigator](/bindings/navigator/Navigator.go)

**navigator/geolocation**

- [geolocation.Geolocation](/bindings/navigator/geolocation/Geolocation.go)
- [geolocation.GeolocationPosition](/bindings/navigator/geolocation/GeolocationPosition.go)
- [geolocation.GeolocationPositionError](/bindings/navigator/geolocation/GeolocationPositionError.go)
- [geolocation.GeolocationPositionOptions](/bindings/navigator/geolocation/GeolocationPositionOptions.go) [1]

**storages**

- [storages.LocalStorage](/bindings/storages/LocalStorage.go)
- [storages.SessionStorage](/bindings/storages/SessionStorage.go)

**timers**

- [timers.ClearInterval](/bindings/timers/ClearInterval.go)
- [timers.ClearTimeout](/bindings/timers/ClearTimeout.go)
- [timers.SetInterval](/bindings/timers/SetInterval.go)
- [timers.SetTimeout](/bindings/timers/SetTimeout.go)

**xhr**

- [xhr.Method](/bindings/xhr/Method.go)
- [xhr.XMLHttpRequest](/bindings/xhr/XMLHttpRequest.go) [2]
- Synchronous [xhr.XMLHttpRequest](/bindings/xhr/XMLHttpRequest_sync.go)

--------

[1] This feature is implemented, but not supported across all Browsers. It is disabled to prevent WebASM runtime errors that are irrecoverable.

[2] This feature is implemented asynchronously and uses a go `chan`. It only works with `tinygo` as a compiler as of now. If your WebASM binary
    hangs when using this, use the synchronous XMLHttpRequest APIs instead.

