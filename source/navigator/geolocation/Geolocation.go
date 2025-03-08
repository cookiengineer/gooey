package geolocation

import "syscall/js"

var Geolocation geolocation

type geolocation struct {
	Value *js.Value `json:"value"`
}

func init() {

	geo := js.Global().Get("window").Get("navigator").Get("geolocation")

	Geolocation = geolocation{
		Value: &geo,
	}

}

func (geolocation *geolocation) GetCurrentPosition(onsuccess func(GeolocationPosition), onerror func(GeolocationPositionError)) {

	wrapped_onsuccess := js.FuncOf(func(this js.Value, args []js.Value) any {

		if len(args) == 1 {

			position := ToGeolocationPosition(args[0])
			onsuccess(position)

		}

		return nil

	})

	wrapped_onerror := js.FuncOf(func(this js.Value, args []js.Value) any {

		if len(args) == 1 {

			wrapped_error := args[0]

			if !wrapped_error.IsUndefined() && !wrapped_error.IsNull() {

				code := wrapped_error.Get("code")

				if !code.IsUndefined() && !code.IsNull() {
					onerror(GeolocationPositionError(code.Int()))
				} else {
					onerror(GeolocationPositionErrorUnknown)
				}

			} else {
				onerror(GeolocationPositionErrorUnknown)
			}

		}

		return nil

	})

	// GeolocationOptions don't work in all Browsers
	wrapped_options := js.ValueOf(nil)

	geolocation.Value.Call("getCurrentPosition", wrapped_onsuccess, wrapped_onerror, wrapped_options)

}

func (geolocation *geolocation) WatchPosition(onsuccess func(GeolocationPosition), onerror func(GeolocationPositionError)) int {

	var handler_id int = -1

	wrapped_onsuccess := js.FuncOf(func(this js.Value, args []js.Value) any {

		if len(args) == 1 {

			position := ToGeolocationPosition(args[0])
			onsuccess(position)

		}

		return nil

	})

	wrapped_onerror := js.FuncOf(func(this js.Value, args []js.Value) any {

		if len(args) == 1 {

			wrapped_error := args[0]

			if !wrapped_error.IsUndefined() && !wrapped_error.IsNull() {

				code := wrapped_error.Get("code")

				if !code.IsUndefined() && !code.IsNull() {
					onerror(GeolocationPositionError(code.Int()))
				} else {
					onerror(GeolocationPositionErrorUnknown)
				}

			} else {
				onerror(GeolocationPositionErrorUnknown)
			}

		}

		return nil

	})

	// GeolocationOptions don't work in all Browsers
	wrapped_options := js.ValueOf(nil)
	wrapped_handler := geolocation.Value.Call("watchPosition", wrapped_onsuccess, wrapped_onerror, wrapped_options)

	if !wrapped_handler.IsUndefined() && !wrapped_handler.IsNull() {
		handler_id = wrapped_handler.Int()
	}

	return handler_id

}
