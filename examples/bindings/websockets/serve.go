package main

import "github.com/cookiengineer/gowebsocket"
import "log"
import "net/http"
import "os"

func main() {

	go func() {

		// WebSocket Server
		server := &gowebsocket.Server{
			Addr:     ":3001",
			Handler: func(websocket *gowebsocket.WebSocket) {

				log.Println("-> Client connected")

				websocket.OnMessage = func(message []byte) {
					websocket.Send([]byte("Hello from server!"))
				}

				websocket.OnClose = func() {
					log.Println("-> Client disconnected")
				}

			},
			TLSConfig: nil,
			ErrorLog:  nil,
		}

		server.Listen()

	}()

	// HTTP Server
	fsys := os.DirFS("public")
	fsrv := http.FileServer(http.FS(fsys))

	http.Handle("/", fsrv)

	log.Println("Listening on http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
