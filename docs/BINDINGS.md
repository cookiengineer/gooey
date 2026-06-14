
# Gooey Bindings

## Implementation Status

This table documents the implementation status of the Gooey [bindings](/bindings)
package, and whether or not the bindings themselves are considered stable. If a
package is documented more thoroughly and has examples in the API documentation,
it is considered stable for public use.

As gooey will run via WebASM in Web Browsers, no assumption is being made about
compatibility of Web Browser Engines. Use at your own risk, test at your own peril.

| Stable? | Implementation                                                                                                             | API Docs                                                                                                            |
|:-------:|:---------------------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------:|
| yes     | [bindings#Screen](/bindings/Screen.go)                                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings#Screen)                                           |
| yes     | [bindings#ScreenOrientation](/bindings/ScreenOrientation.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings#ScreenOrientation)                                |
| yes     | [bindings#Window](/bindings/Window.go)                                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings#Window)                                           |
| yes     | [bindings/animations#CancelAnimationFrame](/bindings/animations/CancelAnimationFrame.go)                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/animations#CancelAnimationFrame)                  |
| yes     | [bindings/animations#RequestAnimationFrame](/bindings/animations/RequestAnimationFrame.go)                                 | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/animations#RequestAnimationFrame)                 |
| TODO    | bindings/canvas2d | |
| yes     | [bindings/console#Console](/bindings/console/Console.go)                                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/console#Console)                                  |
| yes     | [bindings/cookiestore#Cookie](/bindings/cookiestore/Cookie.go)                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#Cookie)                               |
| yes     | [bindings/cookiestore#CookieStore](/bindings/cookiestore/CookieStore.go)                                               [1] | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#CookieStore)                          |
| yes     | [bindings/cookiestore#DeleteOptions](/bindings/cookiestore/DeleteOptions.go)                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#DeleteOptions)                        |
| yes     | [bindings/cookiestore#GetOptions](/bindings/cookiestore/GetOptions.go)                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#GetOptions)                           |
| yes     | [bindings/cookiestore#SameSite](/bindings/cookiestore/SameSite.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#SameSite)                             |
| yes     | [bindings/cookiestore#SetOptions](/bindings/cookiestore/SetOptions.go)                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#SetOptions)                           |
| TODO    | bindings/crypto | |
| TODO    | bindings/crypto/aescbc | |
| TODO    | bindings/crypto/aesctr | |
| TODO    | bindings/crypto/aesgcm | |
| TODO    | bindings/dom | |
| yes     | [bindings/encoding#TextDecoder](/bindings/encoding/TextDecoder.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/encoding#TextDecoder)                             |
| yes     | [bindings/encoding#TextDecoderOptions](/bindings/encoding/TextDecoderOptions.go)                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/encoding#TextDecoderOptions)                      |
| yes     | [bindings/encoding#TextEncoder](/bindings/encoding/TextEncoder.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/encoding#TextEncoder)                             |
| yes     | [bindings/fetch#Cache](/bindings/fetch/Cache.go)                                                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Cache)                                      |
| yes     | [bindings/fetch#Credentials](/bindings/fetch/Credentials.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Credentials)                                |
| yes     | [bindings/fetch#Fetch](/bindings/fetch/Fetch.go)                                                                       [1] | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Fetch)                                      |
| yes     | [bindings/fetch#Headers](/bindings/fetch/Headers.go)                                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Headers)                                    |
| yes     | [bindings/fetch#Method](/bindings/fetch/Method.go)                                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Method)                                     |
| yes     | [bindings/fetch#Mode](/bindings/fetch/Mode.go)                                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Mode)                                       |
| yes     | [bindings/fetch#Redirect](/bindings/fetch/Redirect.go)                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Redirect)                                   |
| yes     | [bindings/fetch#ReferrerPolicy](/bindings/fetch/ReferrerPolicy.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#ReferrerPolicy)                             |
| yes     | [bindings/fetch#RequestInit](/bindings/fetch/RequestInit.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#RequestInit)                                |
| yes     | [bindings/fetch#Response](/bindings/fetch/Response.go)                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Response)                                   |
| TODO    | bindings/history | |
| yes     | [bindings/location#Location](/bindings/location/Location.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/location#Location)                                |
| yes     | [bindings/navigator#Navigator](/bindings/navigator/Navigator.go)                                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator#Navigator)                              |
| yes     | [bindings/navigator/geolocation#Geolocation](/bindings/navigator/geolocation/Geolocation.go)                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator/geolocation#Geolocation)                |
| yes     | [bindings/navigator/geolocation#GeolocationPosition](/bindings/navigator/geolocation/GeolocationPosition.go)               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator/geolocation#GeolocationPosition)        |
| yes     | [bindings/navigator/geolocation#GeolocationPositionError](/bindings/navigator/geolocation/GeolocationPositionError.go)     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator/geolocation#GeolocationPositionError)   |
| yes     | [bindings/navigator/geolocation#GeolocationPositionOptions](/bindings/navigator/geolocation/GeolocationPositionOptions.go) | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator/geolocation#GeolocationPositionOptions) |
| yes     | [bindings/storages#LocalStorage](/bindings/storages/LocalStorage.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/storages#LocalStorage)                            |
| yes     | [bindings/storages#SessionStorage](/bindings/storages/SessionStorage.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/storages#SessionStorage)                          |
| yes     | [bindings/storages#Storage](/bindings/storages/Storage.go)                                                                 | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/storages#Storage)                                 |
| yes     | [bindings/timers#ClearInterval](/bindings/timers/ClearInterval.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/timers#ClearInterval)                             |
| yes     | [bindings/timers#ClearTimeout](/bindings/timers/ClearTimeout.go)                                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/timers#ClearTimeout)                              |
| yes     | [bindings/timers#SetInterval](/bindings/timers/SetInterval.go)                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/timers#SetInterval)                               |
| yes     | [bindings/timers#SetTimeout](/bindings/timers/SetTimeout.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/timers#SetTimeout)                                |
| yes     | [bindings/xhr#Method](/bindings/xhr/Method.go)                                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/xhr#Method)                                       |
| yes     | [bindings/xhr#XMLHttpRequest](/bindings/xhr/XMLHttpRequest.go)                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/xhr#XMLHttpRequest)                               |

[1] This feature is implemented asynchronously and uses go channels. Use a `go func(){}` around it to prevent deadlocks. See [ERRATA.md](./ERRATA.md) for details.

## Implementation Notes

**fetch**:

- `RequestInit.Duplex` property is not supported.
- `RequestInit.Referrer` property has to be a `string` due to arbitrary URL values.

**xhr**:

- `XMLHttpRequest.Upload` monitoring property is not supported.

--------

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

- [dom.Document](/bindings/dom/Document.go)
- [dom.Element](/bindings/dom/Element.go)
- [dom.Event](/bindings/dom/Event.go)
- [dom.EventListener](/bindings/dom/EventListener.go)
- [dom.EventPhase](/bindings/dom/EventPhase.go)
- [dom.EventType](/bindings/dom/EventType.go)

**history**

- [history.EventListener](/bindings/history/EventListener.go)
- [history.EventPhase](/bindings/history/EventPhase.go)
- [history.EventType](/bindings/history/EventType.go)
- [history.History](/bindings/history/History.go)
- [history.HistoryState](/bindings/history/HistoryState.go)
- [history.PopStateEvent](/bindings/history/PopStateEvent.go)

--------

