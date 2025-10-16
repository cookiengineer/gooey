
# Components

## cookiestore Example

## components/content

- [ ] Implement `content.Fieldset` setter methods
- [ ] Implement `content.LineChart` setter methods
- [ ] Implement `content.PieChart` setter methods
- [ ] Implement `content.Table` setter methods

## components/interfaces

- [ ] Decide on an interface for a Schema of `N Data` (`map[string]any`) instances
- [ ] Decide on an interface for a Schema of `1 Data` (`map[string]any`) instances
- [ ] Decide how to integrate custom `String()` methods on a Schema instance, which has to allow
      the generation of HTML, the generation of strings themselves and be integratable into custom
      components. Maybe it makes sense to have a Renderable or Stringifiable interface that is used
      by the components?

## app/Main

- [ ] Integrate the History API in the fallback case, instead of using `location.Location.Replace()`

## components/content/BarChart

- [ ] Use a `<figure>` element as a wrapper
- [ ] chart.Data is a data.Data instance
- [ ] Labels []string
- [ ] Properties []string

## components/content/RadialChart

- [ ] chart.Dataset is data.Dataset instance
- [ ] Labels []string
- [ ] Properties []string

## components/ui/Choices

- [ ] Needs a separate `<div>` element surrounding the `<input type="radio">` elements
- [ ] Root element is the `div`, not any of the input radio elements.

## components/ui/Progress

- [ ] Use `progress` element
- [ ] Min/Max/Step `int` Properties

## components/ui/Login

- [ ] WebAuthN integrated login components
- [ ] 2FA integrated login components
- [ ] OpenAuth has best development experience
- [ ] https://openauth.js.org/
- [ ] Don't mix userdata with authentication
- [ ] Implement support for keycloak service as example?


# Bindings

## bindings/canvas2d

- [ ] `context.RoundRect()`

Canvas Images:

- [ ] canvas2d/Image
- [ ] canvas2d/ImageBitmap
- [ ] canvas2d/ImageData
- [ ] `canvas2d/Canvas` needs `ToDataURL()` and `ToBlob()` support
- [ ] `context.DrawImage()`

Canvas Paths:

- [ ] `Context.ClipPath(path Path2D)` to call `context.clip(path)` API
- [ ] See also [Path2D](https://developer.mozilla.org/en-US/docs/Web/API/Path2D)
- [ ] `context.IsPointInPath()`
- [ ] `context.IsPointInStroke()`

Canvas Styles:

- [ ] `context.SetFillStyleColor()`
- [ ] `context.SetFillStylePattern()`
- [ ] `context.SetFillStyleGradient()`

- [ ] `context.SetStrokeStyleColor()`
- [ ] `context.SetStrokeStylePattern()`
- [ ] `context.SetStrokeStyleGradient()`

Canvas Gradients:

- [ ] canvas2d/CanvasGradient
- [ ] `context.CreateConicGradient()`
- [ ] `context.CreateLinearGradient()`
- [ ] `context.CreateRadialGradient()`

Canvas Patterns:

- [ ] canvas2d/CanvasPattern
- [ ] `context.CreatePattern()`

Canvas Matrixes:

- [ ] canvas2d/DOMMatrix
- [ ] `context.SetTransformMatrix(matrix DOMMatrix)`

## bindings/css

- [ ] `css/FontFace`
- [ ] How to implement font descriptors?
- [ ] How to implement unicode ranges?

- [ ] `css/Matrix`
- [ ] `css/Length`

## bindings/crypto/aesgcm

- [ ] Maybe there's a use case for AES-GCM's `additionalData` and `tagLength` parameters of
      the `AesGcmParams` object. If there is, the `aesgcm.Encrypt()` method needs to change.

## bindings/crypto/pbkdf2

- [ ] Web API only supports `DeriveKey()` method.
- [ ] See also [Pbkdf2Params](https://developer.mozilla.org/en-US/docs/Web/API/Pbkdf2Params)

## bindings/crypto/ecdh

- [ ] ECDH can only be used to `DeriveKey()`
- [ ] See also [EcKeyGenparams](https://developer.mozilla.org/en-US/docs/Web/API/EcKeyGenParams)
- [ ] See also [deriveKey](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/deriveKey)

## bindings/crypto/ecdsa

- [ ] ECDSA can only be used to `Sign()`
- [ ] See also [EcKeyGenparams](https://developer.mozilla.org/en-US/docs/Web/API/EcKeyGenParams)
- [ ] See also [sign](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/sign)

## bindings/navigator

[Navigator Object](https://html.spec.whatwg.org/multipage/system-state.html#the-navigator-object):

- [ ] `OnLine` property might change
- [ ] `DoNotTrack` property might change
- [ ] `CookieEnabled` property might change
- [ ] See also [Clipboard API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/clipboard)
- [ ] See also [Credentials API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/credentials)
- [ ] See also [Keyboard API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/keyboard)
- [ ] See also [Permissions API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/permissions)
- [ ] See also [Storage API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/storage)
- [ ] See also [VirtualKeyBoard API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/virtualKeyboard)
- [ ] See also [Vibration API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/vibrate)

