//go:build wasm

// Package console implements the Console Living Standard specified by WHATWG and tries to be as
// compatible as possible to the specification.
//
// See also: https://console.spec.whatwg.org/
//
// Example usage:
//
//	import "github.com/cookiengineer/gooey/bindings/console"
//
//	console1 := console.GetConsole()
//	console1.Log(1337)
//	console1.Debug("foo bar")
package console

import "encoding/json"
import "syscall/js"

var global_console *Console

func init() {
	global_console = GetConsole()
}

type Console struct {
	Value *js.Value `json:"value"`
}

func GetConsole() *Console {

	if global_console != nil {

		return global_console

	} else {

		value := js.Global().Get("window").Get("console")

		console := Console{
			Value: &value,
		}

		return &console

	}

}

func (console *Console) Assert(condition bool, raw any) {

	switch raw.(type) {
	case js.Value:
		console.Value.Call("assert", js.ValueOf(condition), raw)

	case error:
		value := js.ValueOf(raw.(error).Error())
		console.Value.Call("assert", js.ValueOf(condition), value)

	case []byte:
		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))
		js.CopyBytesToJS(value, bytes)
		console.Value.Call("assert", js.ValueOf(condition), value)

	case string:
		value := js.ValueOf(raw.(string))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case int:
		value := js.ValueOf(raw.(int))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case int8:
		value := js.ValueOf(raw.(int8))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case int16:
		value := js.ValueOf(raw.(int16))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case int32:
		value := js.ValueOf(raw.(int32))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case int64:
		value := js.ValueOf(raw.(int64))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case uint:
		value := js.ValueOf(raw.(uint))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case uint8:
		value := js.ValueOf(raw.(uint8))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case uint16:
		value := js.ValueOf(raw.(uint16))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case uint32:
		value := js.ValueOf(raw.(uint32))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case uint64:
		value := js.ValueOf(raw.(uint64))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case float32:
		value := js.ValueOf(raw.(float32))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case float64:
		value := js.ValueOf(raw.(float64))
		console.Value.Call("assert", js.ValueOf(condition), value)

	case any:
		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {
			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)
			console.Value.Call("assert", js.ValueOf(condition), object)
		}

	}

}

func (console *Console) Clear() {
	console.Value.Call("clear")
}

func (console *Console) Count(label string) {
	console.Value.Call("count", js.ValueOf(label))
}

func (console *Console) CountReset(label string) {
	console.Value.Call("countReset", js.ValueOf(label))
}

func (console *Console) Debug(raw any) {

	switch raw.(type) {
	case js.Value:
		console.Value.Call("debug", raw)

	case error:
		value := js.ValueOf(raw.(error).Error())
		console.Value.Call("debug", value)

	case []byte:
		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))
		js.CopyBytesToJS(value, bytes)
		console.Value.Call("debug", value)

	case string:
		value := js.ValueOf(raw.(string))
		console.Value.Call("debug", value)

	case int:
		value := js.ValueOf(raw.(int))
		console.Value.Call("debug", value)

	case int8:
		value := js.ValueOf(raw.(int8))
		console.Value.Call("debug", value)

	case int16:
		value := js.ValueOf(raw.(int16))
		console.Value.Call("debug", value)

	case int32:
		value := js.ValueOf(raw.(int32))
		console.Value.Call("debug", value)

	case int64:
		value := js.ValueOf(raw.(int64))
		console.Value.Call("debug", value)

	case uint:
		value := js.ValueOf(raw.(uint))
		console.Value.Call("debug", value)

	case uint8:
		value := js.ValueOf(raw.(uint8))
		console.Value.Call("debug", value)

	case uint16:
		value := js.ValueOf(raw.(uint16))
		console.Value.Call("debug", value)

	case uint32:
		value := js.ValueOf(raw.(uint32))
		console.Value.Call("debug", value)

	case uint64:
		value := js.ValueOf(raw.(uint64))
		console.Value.Call("debug", value)

	case float32:
		value := js.ValueOf(raw.(float32))
		console.Value.Call("debug", value)

	case float64:
		value := js.ValueOf(raw.(float64))
		console.Value.Call("debug", value)

	case any:
		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {
			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)
			console.Value.Call("debug", object)
		}

	}

}

func (console *Console) Error(raw any) {

	switch raw.(type) {
	case js.Value:
		console.Value.Call("error", raw)

	case error:
		value := js.ValueOf(raw.(error).Error())
		console.Value.Call("error", value)

	case []byte:
		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))
		js.CopyBytesToJS(value, bytes)
		console.Value.Call("error", value)

	case string:
		value := js.ValueOf(raw.(string))
		console.Value.Call("error", value)

	case int:
		value := js.ValueOf(raw.(int))
		console.Value.Call("error", value)

	case int8:
		value := js.ValueOf(raw.(int8))
		console.Value.Call("error", value)

	case int16:
		value := js.ValueOf(raw.(int16))
		console.Value.Call("error", value)

	case int32:
		value := js.ValueOf(raw.(int32))
		console.Value.Call("error", value)

	case int64:
		value := js.ValueOf(raw.(int64))
		console.Value.Call("error", value)

	case uint:
		value := js.ValueOf(raw.(uint))
		console.Value.Call("error", value)

	case uint8:
		value := js.ValueOf(raw.(uint8))
		console.Value.Call("error", value)

	case uint16:
		value := js.ValueOf(raw.(uint16))
		console.Value.Call("error", value)

	case uint32:
		value := js.ValueOf(raw.(uint32))
		console.Value.Call("error", value)

	case uint64:
		value := js.ValueOf(raw.(uint64))
		console.Value.Call("error", value)

	case float32:
		value := js.ValueOf(raw.(float32))
		console.Value.Call("error", value)

	case float64:
		value := js.ValueOf(raw.(float64))
		console.Value.Call("error", value)

	case any:
		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {
			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)
			console.Value.Call("error", object)
		}

	}

}

func (console *Console) Group(label string) {
	console.Value.Call("group", js.ValueOf(label))
}

func (console *Console) GroupCollapsed(label string) {
	console.Value.Call("groupCollapsed", js.ValueOf(label))
}

func (console *Console) GroupEnd() {
	console.Value.Call("groupEnd")
}

func (console *Console) Info(raw any) {

	switch raw.(type) {
	case js.Value:
		console.Value.Call("info", raw)

	case error:
		value := js.ValueOf(raw.(error).Error())
		console.Value.Call("info", value)

	case []byte:
		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))
		js.CopyBytesToJS(value, bytes)
		console.Value.Call("info", value)

	case string:
		value := js.ValueOf(raw.(string))
		console.Value.Call("info", value)

	case int:
		value := js.ValueOf(raw.(int))
		console.Value.Call("info", value)

	case int8:
		value := js.ValueOf(raw.(int8))
		console.Value.Call("info", value)

	case int16:
		value := js.ValueOf(raw.(int16))
		console.Value.Call("info", value)

	case int32:
		value := js.ValueOf(raw.(int32))
		console.Value.Call("info", value)

	case int64:
		value := js.ValueOf(raw.(int64))
		console.Value.Call("info", value)

	case uint:
		value := js.ValueOf(raw.(uint))
		console.Value.Call("info", value)

	case uint8:
		value := js.ValueOf(raw.(uint8))
		console.Value.Call("info", value)

	case uint16:
		value := js.ValueOf(raw.(uint16))
		console.Value.Call("info", value)

	case uint32:
		value := js.ValueOf(raw.(uint32))
		console.Value.Call("info", value)

	case uint64:
		value := js.ValueOf(raw.(uint64))
		console.Value.Call("info", value)

	case float32:
		value := js.ValueOf(raw.(float32))
		console.Value.Call("info", value)

	case float64:
		value := js.ValueOf(raw.(float64))
		console.Value.Call("info", value)

	case any:
		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {
			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)
			console.Value.Call("info", object)
		}

	}

}

func (console *Console) Log(raw any) {

	switch raw.(type) {
	case js.Value:
		console.Value.Call("log", raw)

	case error:
		value := js.ValueOf(raw.(error).Error())
		console.Value.Call("log", value)

	case []byte:
		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))
		js.CopyBytesToJS(value, bytes)
		console.Value.Call("log", value)

	case string:
		value := js.ValueOf(raw.(string))
		console.Value.Call("log", value)

	case int:
		value := js.ValueOf(raw.(int))
		console.Value.Call("log", value)

	case int8:
		value := js.ValueOf(raw.(int8))
		console.Value.Call("log", value)

	case int16:
		value := js.ValueOf(raw.(int16))
		console.Value.Call("log", value)

	case int32:
		value := js.ValueOf(raw.(int32))
		console.Value.Call("log", value)

	case int64:
		value := js.ValueOf(raw.(int64))
		console.Value.Call("log", value)

	case uint:
		value := js.ValueOf(raw.(uint))
		console.Value.Call("log", value)

	case uint8:
		value := js.ValueOf(raw.(uint8))
		console.Value.Call("log", value)

	case uint16:
		value := js.ValueOf(raw.(uint16))
		console.Value.Call("log", value)

	case uint32:
		value := js.ValueOf(raw.(uint32))
		console.Value.Call("log", value)

	case uint64:
		value := js.ValueOf(raw.(uint64))
		console.Value.Call("log", value)

	case float32:
		value := js.ValueOf(raw.(float32))
		console.Value.Call("log", value)

	case float64:
		value := js.ValueOf(raw.(float64))
		console.Value.Call("log", value)

	case any:
		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {
			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)
			console.Value.Call("log", object)
		}

	}

}

func (console *Console) Time(label string) {
	console.Value.Call("time", js.ValueOf(label))
}

func (console *Console) TimeEnd(label string) {
	console.Value.Call("timeEnd", js.ValueOf(label))
}

func (console *Console) TimeLog(label string, raw any) {

	switch raw.(type) {
	case js.Value:
		console.Value.Call("timeLog", js.ValueOf(label), raw)

	case error:
		value := js.ValueOf(raw.(error).Error())
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case []byte:
		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))
		js.CopyBytesToJS(value, bytes)
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case string:
		value := js.ValueOf(raw.(string))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case int:
		value := js.ValueOf(raw.(int))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case int8:
		value := js.ValueOf(raw.(int8))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case int16:
		value := js.ValueOf(raw.(int16))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case int32:
		value := js.ValueOf(raw.(int32))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case int64:
		value := js.ValueOf(raw.(int64))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case uint:
		value := js.ValueOf(raw.(uint))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case uint8:
		value := js.ValueOf(raw.(uint8))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case uint16:
		value := js.ValueOf(raw.(uint16))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case uint32:
		value := js.ValueOf(raw.(uint32))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case uint64:
		value := js.ValueOf(raw.(uint64))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case float32:
		value := js.ValueOf(raw.(float32))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case float64:
		value := js.ValueOf(raw.(float64))
		console.Value.Call("timeLog", js.ValueOf(label), value)

	case any:
		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {
			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)
			console.Value.Call("timeLog", js.ValueOf(label), object)
		}

	}

}

func (console *Console) Trace(raw any) {

	switch raw.(type) {
	case js.Value:
		console.Value.Call("trace", raw)

	case error:
		value := js.ValueOf(raw.(error).Error())
		console.Value.Call("trace", value)

	case []byte:
		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))
		js.CopyBytesToJS(value, bytes)
		console.Value.Call("trace", value)

	case string:
		value := js.ValueOf(raw.(string))
		console.Value.Call("trace", value)

	case int:
		value := js.ValueOf(raw.(int))
		console.Value.Call("trace", value)

	case int8:
		value := js.ValueOf(raw.(int8))
		console.Value.Call("trace", value)

	case int16:
		value := js.ValueOf(raw.(int16))
		console.Value.Call("trace", value)

	case int32:
		value := js.ValueOf(raw.(int32))
		console.Value.Call("trace", value)

	case int64:
		value := js.ValueOf(raw.(int64))
		console.Value.Call("trace", value)

	case uint:
		value := js.ValueOf(raw.(uint))
		console.Value.Call("trace", value)

	case uint8:
		value := js.ValueOf(raw.(uint8))
		console.Value.Call("trace", value)

	case uint16:
		value := js.ValueOf(raw.(uint16))
		console.Value.Call("trace", value)

	case uint32:
		value := js.ValueOf(raw.(uint32))
		console.Value.Call("trace", value)

	case uint64:
		value := js.ValueOf(raw.(uint64))
		console.Value.Call("trace", value)

	case float32:
		value := js.ValueOf(raw.(float32))
		console.Value.Call("trace", value)

	case float64:
		value := js.ValueOf(raw.(float64))
		console.Value.Call("trace", value)

	case any:
		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {
			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)
			console.Value.Call("trace", object)
		}

	}

}

func (console *Console) Warn(raw any) {

	switch raw.(type) {
	case js.Value:
		console.Value.Call("warn", raw)

	case error:
		value := js.ValueOf(raw.(error).Error())
		console.Value.Call("warn", value)

	case []byte:
		bytes := raw.([]byte)
		value := js.Global().Get("Uint8Array").New(len(bytes))
		js.CopyBytesToJS(value, bytes)
		console.Value.Call("warn", value)

	case string:
		value := js.ValueOf(raw.(string))
		console.Value.Call("warn", value)

	case int:
		value := js.ValueOf(raw.(int))
		console.Value.Call("warn", value)

	case int8:
		value := js.ValueOf(raw.(int8))
		console.Value.Call("warn", value)

	case int16:
		value := js.ValueOf(raw.(int16))
		console.Value.Call("warn", value)

	case int32:
		value := js.ValueOf(raw.(int32))
		console.Value.Call("warn", value)

	case int64:
		value := js.ValueOf(raw.(int64))
		console.Value.Call("warn", value)

	case uint:
		value := js.ValueOf(raw.(uint))
		console.Value.Call("warn", value)

	case uint8:
		value := js.ValueOf(raw.(uint8))
		console.Value.Call("warn", value)

	case uint16:
		value := js.ValueOf(raw.(uint16))
		console.Value.Call("warn", value)

	case uint32:
		value := js.ValueOf(raw.(uint32))
		console.Value.Call("warn", value)

	case uint64:
		value := js.ValueOf(raw.(uint64))
		console.Value.Call("warn", value)

	case float32:
		value := js.ValueOf(raw.(float32))
		console.Value.Call("warn", value)

	case float64:
		value := js.ValueOf(raw.(float64))
		console.Value.Call("warn", value)

	case any:
		buffer, err := json.MarshalIndent(raw, "", "\t")

		if err == nil {
			value := js.ValueOf(string(buffer))
			object := js.Global().Get("JSON").Call("parse", value)
			console.Value.Call("warn", object)
		}

	}

}
