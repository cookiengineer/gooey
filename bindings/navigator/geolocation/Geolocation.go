//go:build wasm

package geolocation

import "syscall/js"

var global_geolocation *Geolocation

func init() {
	global_geolocation = GetGeolocation()
}

type Geolocation struct {
	Value *js.Value `json:"value"`
}

// Returns the global Geolocation instance.
func GetGeolocation() *Geolocation {

	if global_geolocation != nil {

		return global_geolocation

	} else {

		geo := js.Global().Get("window").Get("navigator").Get("geolocation")

		geolocation := Geolocation{
			Value: &geo,
		}

		return &geolocation

	}

}

// Clears the watch handler with an identifier obtained by handler_id := WatchPosition(func(GeolocationPostion){}, func(GeolocationError){})
func (geolocation *Geolocation) ClearWatch(handler_id int) {

	wrapped_id := js.ValueOf(handler_id)
	geolocation.Value.Call("clearWatch", wrapped_id)

}

// Gets the current position of the device.
func (geolocation *Geolocation) GetCurrentPosition(onsuccess func(GeolocationPosition), onerror func(GeolocationPositionError)) {

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

// Watches the current position of the device. Returns the identifier used for ClearWatch(handler_id).
func (geolocation *Geolocation) WatchPosition(onsuccess func(GeolocationPosition), onerror func(GeolocationPositionError)) int {

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
