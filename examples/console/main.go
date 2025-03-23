package main

import "github.com/cookiengineer/gooey/bindings/console"
import "errors"
import "syscall/js"
import "time"

type Message struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Bytes []byte `json:"bytes"`
}

func main() {

	// console.Log supports go-native data types
	bytes := []byte{0x01,0x03,0x03,0x07}

	console.Group("bytes")
	console.Log(bytes)
	console.GroupEnd("bytes")

	// console supports errors
	err := errors.New("This is an error with a custom message")

	console.Group("error")
	console.Error(err)
	console.GroupEnd("error")

	// console supports js.Value instances
	js_value := js.Global().Get("Uint8Array").New(4)
	js_value.SetIndex(0, 0x01)
	js_value.SetIndex(1, 0x03)
	js_value.SetIndex(2, 0x03)
	js_value.SetIndex(3, 0x07)

	console.Group("js.Value")
	console.Log(js_value)
	console.GroupEnd("js.Value")

	// console supports struct instances
	message := Message{
		Id:    1337,
		Name:  "cookiengineer",
		Bytes: []byte("This is an example"),
	}

	console.Group("structs")
	console.Log(message)
	console.GroupEnd("structs")

	console.Log("This is a Log")
	console.Info("This is an Information")
	console.Warn("This is a Warning")
	console.Error("This is an Error")

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
