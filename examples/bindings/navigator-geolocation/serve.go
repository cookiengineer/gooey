package main

import "fmt"
import "log"
import "net/http"
import "os"

func main() {

	fsys := os.DirFS("public")
	fsrv := http.FileServer(http.FS(fsys))

	http.Handle("/", fsrv)

	fmt.Println("Listening on http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
