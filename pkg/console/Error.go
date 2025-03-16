//go:build wasm

package console

import "encoding/json"
import "syscall/js"

func Error(raw any) {

	switch raw.(type) {
	case js.Value:

		js.Global().Get("console").Call("error", raw)

	case error:

		value := js.ValueOf(raw.(error).Error())
		js.Global().Get("console").Call("error", value)

	case []byte:

		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))

		js.CopyBytesToJS(value, bytes)
		js.Global().Get("console").Call("error", value)

	case string:

		value := js.ValueOf(raw.(string))
		js.Global().Get("console").Call("error", value)

	case int:
		value := js.ValueOf(raw.(int))
		js.Global().Get("console").Call("error", value)

	case int8:

		value := js.ValueOf(raw.(int8))
		js.Global().Get("console").Call("error", value)

	case int16:

		value := js.ValueOf(raw.(int16))
		js.Global().Get("console").Call("error", value)

	case int32:

		value := js.ValueOf(raw.(int32))
		js.Global().Get("console").Call("error", value)

	case int64:

		value := js.ValueOf(raw.(int64))
		js.Global().Get("console").Call("error", value)

	case uint:

		value := js.ValueOf(raw.(uint))
		js.Global().Get("console").Call("error", value)

	case uint8:

		value := js.ValueOf(raw.(uint8))
		js.Global().Get("console").Call("error", value)

	case uint16:

		value := js.ValueOf(raw.(uint16))
		js.Global().Get("console").Call("error", value)

	case uint32:

		value := js.ValueOf(raw.(uint32))
		js.Global().Get("console").Call("error", value)

	case uint64:

		value := js.ValueOf(raw.(uint64))
		js.Global().Get("console").Call("error", value)

	case float32:

		value := js.ValueOf(raw.(float32))
		js.Global().Get("console").Call("error", value)

	case float64:

		value := js.ValueOf(raw.(float64))
		js.Global().Get("console").Call("error", value)

	case any:

		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {

			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)

			js.Global().Get("console").Call("error", object)

		}

	}

}
