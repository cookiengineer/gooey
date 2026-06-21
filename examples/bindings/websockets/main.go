package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/websockets"
import "fmt"
import "time"

func main() {

	console   := console.GetConsole()
	websocket := websockets.NewWebSocket("ws://localhost:3001/upgrade-me")

	websocket.AddEventListener(websockets.EventTypeOpen, websockets.ToEventListener(func(event *websockets.Event) {
		console.Log("Open Event")
	}))

	websocket.AddEventListener(websockets.EventTypeMessage, websockets.ToEventListener(func(event *websockets.Event) {
		console.Log(fmt.Sprintf("Message Event: %s", string(event.Data)))
	}))

	websocket.AddEventListener(websockets.EventTypeClose, websockets.ToEventListener(func(event *websockets.Event) {
		console.Log("Close Event")
	}))

	go func() {
		time.Sleep(1 * time.Second)
		websocket.Send([]byte("Hello, world!"))
	}()


	for true {

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
