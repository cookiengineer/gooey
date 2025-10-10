package main

import "fmt"
import "log"
import "net/http"
import "os"

func main() {

	fsys := os.DirFS("public")
	fsrv := http.FileServer(http.FS(fsys))

	http.Handle("/", fsrv)

	http.HandleFunc("/api/test", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			response.Write([]byte("{\"message\": \"success\"}"))

		} else if request.Method == http.MethodOptions {

			origin := request.URL.Scheme + "://" + request.URL.Host

			response.Header().Set("Access-Control-Allow-Origin", origin)
			response.WriteHeader(http.StatusOK)
			response.Write([]byte(""))

		}

	})

	fmt.Println("Listening on http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
