
# Components

## Oauth Components

- [ ] WebAuthN integrated login components
- [ ] 2FA integrated login components
- [ ] OpenAuth has best development experience
- [ ] https://openauth.js.org/
- [ ] Don't mix userdata with authentication
- [ ] Implement support for keycloak service as example?

## Remix and tanstack cookie handling

## App Router? Server-Side?

- [ ] Provide a server middleware for routing, auth and server-side rendering
- [ ] How to map server-side routes?
- [ ] Should ideally reuse the app.Main and app.View based workflow, meaning
      there should be an implementation of the virtual DOM that's independent of
      the `syscall/js` interfaces.
- [ ] Server-side auth middleware should use `context`

- [ ] OpenAPI on-the-fly generation?

## content/Fieldset

- [ ] Implement a `Reset()` method that resets all fields

## content/Article

- [ ] Create a `content.Article` component

## content/Table

- [ ] Create a `content.Table` component

## app/View

- [ ] Integrate the History API in the fallback case, instead of using `location.Location.Replace()`

## layout/Dialog

- [ ] `SetPrimaryAction(label string, action string)` method
- [ ] `SetSecondaryAction(label string, action string)` method
- [ ] Footer property should be a `layout.Footer`

- [ ] `Dialog.Parse()` should parse fieldset, table, and other content elements

## ui.Choices

- [ ] Needs a separate `<div>` element surrounding the `<input type="radio">` elements
- [ ] Root element is the `div`, not any of the input radio elements.


## app.View

- [ ] Element needs to be a `<section>` element
- [ ] View needs a `Layout` property (flex, grid, flow)

## Layout Components

- [ ] `app.View` CSS needs to be implemented in `/design/components/app/View.css`

# Bindings

## Encoding API

[Encoders and Decoders](https://encoding.spec.whatwg.org/#encoders-and-decoders):

- [ ] encoding/TextDecoder
- [ ] encoding/TextEncoder


## CookieStore API

[CookieStore Object](https://developer.mozilla.org/en-US/docs/Web/API/CookieStore)


## Navigator API

[Navigator Object](https://html.spec.whatwg.org/multipage/system-state.html#the-navigator-object):

- [ ] OnLine property might change
- [ ] DoNotTrack property might change
- [ ] CookieEnabled property might change

- [Clipboard API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/clipboard)
- [Credentials API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/credentials)
- [Geolocation API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/geolocation)
- [Keyboard API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/keyboard)
- [Permissions API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/permissions)
- [Storage API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/storage)
- [VirtualKeyBoard API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/virtualKeyboard)
- [Vibration API](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/vibrate)


## Canvas API

- [ ] canvas/Canvas
- [ ] canvas/CanvasRenderingContext2D
- [ ] canvas/CanvasGradient
- [ ] canvas/CanvasPattern
- [ ] canvas/ImageBitmap
- [ ] canvas/ImageData
- [ ] canvas/TextMetrics
- [ ] canvas/OffscreenCanvas
- [ ] canvas/OffscreenCanvasRenderingContext2D
- [ ] canvas/ImageBitmapRenderingContext


## Crypto API

[Web Crypto API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Crypto_API)

### AES-GCM

- [ ] Maybe there's a use case for AES-GCM's `additionalData` and `tagLength` parameters of
      the `AesGcmParams` object. If there is, the `aesgcm.Encrypt()` method needs to change.

### AES-KW

- [ ] Workflow might be different, and requires different `structs` to interact with `PBKDF2`

### PBKDF2

- [ ] Only supports `DeriveKey()` method.
- [ ] See also [Pbkdf2Params](https://developer.mozilla.org/en-US/docs/Web/API/Pbkdf2Params)

### ECDH / ECDSA

- [ ] See also [EcKeyGenparams](https://developer.mozilla.org/en-US/docs/Web/API/EcKeyGenParams)
- [ ] ECDSA can only be used to `Sign()`
- [ ] ECDH can only be used to `DeriveKey()`
- [ ] See also [sign](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/sign)
- [ ] See also [deriveKey](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/deriveKey)

### Other

- HMAC
- RSA OAEP
- RSA PSS


## Web Forms

- [ ] xhr/FormData interface?
- [ ] fetch/FormData interface?

Web Form Elements:

- [ ] elements/forms/Button
- [ ] elements/forms/Form (that can generate FormData)
- [ ] elements/forms/Input
- [ ] elements/forms/Option
- [ ] elements/forms/Select
- [ ] elements/forms/Textarea

