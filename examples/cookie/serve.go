package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fsys := os.DirFS("public")
	fmt.Println(fsys)
	fsrv := http.FileServer(http.FS(fsys))

	http.Handle("/", fsrv)

	http.HandleFunc("/api/test", func(response http.ResponseWriter, request *http.Request) {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"message\": \"success\"}"))

	})

	fmt.Println("Listening on http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
