
# Gooey Bindings

## Implementation Status

This table documents the implementation status of the Gooey [bindings](/bindings)
package, and whether or not the bindings themselves are considered stable. If a
package is documented more thoroughly and has examples in the API documentation,
it is considered stable for public use.

As gooey will run via WebASM in Web Browsers, no assumption is being made about
compatibility of Web Browser Engines. Use at your own risk, test at your own peril.

| Stable?  | Implementation                                                                                                             | API Docs                                                                                                            |
|:--------:|:---------------------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------:|
| yes      | [bindings#Screen](/bindings/Screen.go)                                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings#Screen)                                           |
| yes      | [bindings#ScreenOrientation](/bindings/ScreenOrientation.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings#ScreenOrientation)                                |
| yes      | [bindings#Window](/bindings/Window.go)                                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings#Window)                                           |
| yes      | [bindings/animations#CancelAnimationFrame](/bindings/animations/CancelAnimationFrame.go)                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/animations#CancelAnimationFrame)                  |
| yes      | [bindings/animations#RequestAnimationFrame](/bindings/animations/RequestAnimationFrame.go)                                 | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/animations#RequestAnimationFrame)                 |
| **TODO** | bindings/canvas2d | |
| yes      | [bindings/console#Console](/bindings/console/Console.go)                                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/console#Console)                                  |
| yes      | [bindings/cookiestore#Cookie](/bindings/cookiestore/Cookie.go)                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#Cookie)                               |
| yes      | [bindings/cookiestore#CookieStore](/bindings/cookiestore/CookieStore.go)                                               [1] | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#CookieStore)                          |
| yes      | [bindings/cookiestore#DeleteOptions](/bindings/cookiestore/DeleteOptions.go)                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#DeleteOptions)                        |
| yes      | [bindings/cookiestore#GetOptions](/bindings/cookiestore/GetOptions.go)                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#GetOptions)                           |
| yes      | [bindings/cookiestore#SameSite](/bindings/cookiestore/SameSite.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#SameSite)                             |
| yes      | [bindings/cookiestore#SetOptions](/bindings/cookiestore/SetOptions.go)                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/cookiestore#SetOptions)                           |
| yes      | [bindings/crypto#GetRandomValues](/bindings/crypto/GetRandomValues.go)                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto#GetRandomValues)                           |
| yes      | [bindings/crypto#RandomUUID](/bindings/crypto/RandomUUID.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto#RandomUUID)                                |
| yes      | [bindings/crypto/aescbc#CryptoKey](/bindings/crypto/aescbc/CryptoKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aescbc#CryptoKey)                          |
| yes      | [bindings/crypto/aescbc#CryptoKeyType](/bindings/crypto/aescbc/CryptoKeyType.go)                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aescbc#CryptoKeyType)                      |
| yes      | [bindings/crypto/aescbc#Decrypt](/bindings/crypto/aescbc/Decrypt.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aescbc#Decrypt)                            |
| yes      | [bindings/crypto/aescbc#Encrypt](/bindings/crypto/aescbc/Encrypt.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aescbc#Encrypt)                            |
| yes      | [bindings/crypto/aescbc#ExportKey](/bindings/crypto/aescbc/ExportKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aescbc#ExportKey)                          |
| yes      | [bindings/crypto/aescbc#GenerateKey](/bindings/crypto/aescbc/GenerateKey.go)                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aescbc#GenerateKey)                        |
| yes      | [bindings/crypto/aescbc#ImportKey](/bindings/crypto/aescbc/ImportKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aescbc#ImportKey)                          |
| yes      | [bindings/crypto/aesctr#CryptoKey](/bindings/crypto/aesctr/CryptoKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesctr#CryptoKey)                          |
| yes      | [bindings/crypto/aesctr#CryptoKeyType](/bindings/crypto/aesctr/CryptoKeyType.go)                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesctr#CryptoKeyType)                      |
| yes      | [bindings/crypto/aesctr#Decrypt](/bindings/crypto/aesctr/Decrypt.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesctr#Decrypt)                            |
| yes      | [bindings/crypto/aesctr#Encrypt](/bindings/crypto/aesctr/Encrypt.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesctr#Encrypt)                            |
| yes      | [bindings/crypto/aesctr#ExportKey](/bindings/crypto/aesctr/ExportKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesctr#ExportKey)                          |
| yes      | [bindings/crypto/aesctr#GenerateKey](/bindings/crypto/aesctr/GenerateKey.go)                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesctr#GenerateKey)                        |
| yes      | [bindings/crypto/aesctr#ImportKey](/bindings/crypto/aesctr/ImportKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesctr#ImportKey)                          |
| yes      | [bindings/crypto/aesgcm#CryptoKey](/bindings/crypto/aesgcm/CryptoKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesgcm#CryptoKey)                          |
| yes      | [bindings/crypto/aesgcm#CryptoKeyType](/bindings/crypto/aesgcm/CryptoKeyType.go)                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesgcm#CryptoKeyType)                      |
| yes      | [bindings/crypto/aesgcm#Decrypt](/bindings/crypto/aesgcm/Decrypt.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesgcm#Decrypt)                            |
| yes      | [bindings/crypto/aesgcm#Encrypt](/bindings/crypto/aesgcm/Encrypt.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesgcm#Encrypt)                            |
| yes      | [bindings/crypto/aesgcm#ExportKey](/bindings/crypto/aesgcm/ExportKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesgcm#ExportKey)                          |
| yes      | [bindings/crypto/aesgcm#GenerateKey](/bindings/crypto/aesgcm/GenerateKey.go)                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesgcm#GenerateKey)                        |
| yes      | [bindings/crypto/aesgcm#ImportKey](/bindings/crypto/aesgcm/ImportKey.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/crypto/aesgcm#ImportKey)                          |
| yes      | [bindings/dom#Document](/bindings/dom/Document.go)                                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/dom#Document)                                     |
| yes      | [bindings/dom#Element](/bindings/dom/Element.go)                                                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/dom#Element)                                      |
| yes      | [bindings/dom#Event](/bindings/dom/Event.go)                                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/dom#Event)                                        |
| yes      | [bindings/dom#EventListener](/bindings/dom/EventListener.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/dom#EventListener)                                |
| yes      | [bindings/dom#EventPhase](/bindings/dom/EventPhase.go)                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/dom#EventPhase)                                   |
| yes      | [bindings/dom#EventType](/bindings/dom/EventType.go)                                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/dom#EventType)                                    |
| yes      | [bindings/dom#Rect](/bindings/dom/Rect.go)                                                                                 | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/dom#Rect)                                         |
| yes      | [bindings/encoding#TextDecoder](/bindings/encoding/TextDecoder.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/encoding#TextDecoder)                             |
| yes      | [bindings/encoding#TextDecoderOptions](/bindings/encoding/TextDecoderOptions.go)                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/encoding#TextDecoderOptions)                      |
| yes      | [bindings/encoding#TextEncoder](/bindings/encoding/TextEncoder.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/encoding#TextEncoder)                             |
| yes      | [bindings/fetch#Cache](/bindings/fetch/Cache.go)                                                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Cache)                                      |
| yes      | [bindings/fetch#Credentials](/bindings/fetch/Credentials.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Credentials)                                |
| yes      | [bindings/fetch#Fetch](/bindings/fetch/Fetch.go)                                                                       [1] | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Fetch)                                      |
| yes      | [bindings/fetch#Headers](/bindings/fetch/Headers.go)                                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Headers)                                    |
| yes      | [bindings/fetch#Method](/bindings/fetch/Method.go)                                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Method)                                     |
| yes      | [bindings/fetch#Mode](/bindings/fetch/Mode.go)                                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Mode)                                       |
| yes      | [bindings/fetch#Redirect](/bindings/fetch/Redirect.go)                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Redirect)                                   |
| yes      | [bindings/fetch#ReferrerPolicy](/bindings/fetch/ReferrerPolicy.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#ReferrerPolicy)                             |
| yes      | [bindings/fetch#RequestInit](/bindings/fetch/RequestInit.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#RequestInit)                                |
| yes      | [bindings/fetch#Response](/bindings/fetch/Response.go)                                                                     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/fetch#Response)                                   |
| **no**   | [bindings/history#EventListener](/bindings/history/EventListener.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/history#EventListener)                            |
| **no**   | [bindings/history#EventPhase](/bindings/history/EventPhase.go)                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/history#EventPhase)                               |
| **no**   | [bindings/history#EventType](/bindings/history/EventType.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/history#EventType)                                |
| **no**   | [bindings/history#History](/bindings/history/History.go)                                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/history#History)                                  |
| **no**   | [bindings/history#HistoryState](/bindings/history/HistoryState.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/history#HistoryState)                             |
| **no**   | [bindings/history#PopStateEvent](/bindings/history/PopStateEvent.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/history#PopStateEvent)                            |
| yes      | [bindings/location#Location](/bindings/location/Location.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/location#Location)                                |
| yes      | [bindings/navigator#Navigator](/bindings/navigator/Navigator.go)                                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator#Navigator)                              |
| yes      | [bindings/navigator/geolocation#Geolocation](/bindings/navigator/geolocation/Geolocation.go)                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator/geolocation#Geolocation)                |
| yes      | [bindings/navigator/geolocation#GeolocationPosition](/bindings/navigator/geolocation/GeolocationPosition.go)               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator/geolocation#GeolocationPosition)        |
| yes      | [bindings/navigator/geolocation#GeolocationPositionError](/bindings/navigator/geolocation/GeolocationPositionError.go)     | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator/geolocation#GeolocationPositionError)   |
| yes      | [bindings/navigator/geolocation#GeolocationPositionOptions](/bindings/navigator/geolocation/GeolocationPositionOptions.go) | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/navigator/geolocation#GeolocationPositionOptions) |
| yes      | [bindings/storages#LocalStorage](/bindings/storages/LocalStorage.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/storages#LocalStorage)                            |
| yes      | [bindings/storages#SessionStorage](/bindings/storages/SessionStorage.go)                                                   | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/storages#SessionStorage)                          |
| yes      | [bindings/storages#Storage](/bindings/storages/Storage.go)                                                                 | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/storages#Storage)                                 |
| yes      | [bindings/timers#ClearInterval](/bindings/timers/ClearInterval.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/timers#ClearInterval)                             |
| yes      | [bindings/timers#ClearTimeout](/bindings/timers/ClearTimeout.go)                                                           | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/timers#ClearTimeout)                              |
| yes      | [bindings/timers#SetInterval](/bindings/timers/SetInterval.go)                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/timers#SetInterval)                               |
| yes      | [bindings/timers#SetTimeout](/bindings/timers/SetTimeout.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/timers#SetTimeout)                                |
| yes      | [bindings/xhr#Method](/bindings/xhr/Method.go)                                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/xhr#Method)                                       |
| yes      | [bindings/xhr#XMLHttpRequest](/bindings/xhr/XMLHttpRequest.go)                                                             | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/xhr#XMLHttpRequest)                               |
| yes      | [bindings/websockets#Event](/bindings/websockets/Event.go)                                                                 | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/websockets#Event)                                 |
| yes      | [bindings/websockets#EventListener](/bindings/websockets/EventListener.go)                                                 | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/websockets#EventListener)                         |
| yes      | [bindings/websockets#EventType](/bindings/websockets/EventType.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/websockets#EventType)                             |
| yes      | [bindings/websockets#ReadyState](/bindings/websockets/ReadyState.go)                                                       | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/websockets#ReadyState)                            |
| yes      | [bindings/websockets#Status](/bindings/websockets/Status.go)                                                               | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/websockets#Status)                                |
| yes      | [bindings/websockets#WebSocket](/bindings/websockets/WebSocket.go)                                                         | [link](https://pkg.go.dev/github.com/cookiengineer/gooey/bindings/websockets#WebSocket)                             |

[1] This feature is implemented asynchronously and uses go channels. Use a `go func(){}` around it to prevent deadlocks. See [ERRATA.md](./ERRATA.md) for details.

## Implementation Notes

**crypto**:

The Web Crypto API has been split up into separate sub-packages, due to the `algorithm` parameters
object not being implementable as a reusable struct that would make sense.

**fetch**:

- `RequestInit.Duplex` property is not supported.
- `RequestInit.Referrer` property has to be a `string` due to arbitrary URL values.

**xhr**:

- `XMLHttpRequest.Upload` monitoring property is not supported.

