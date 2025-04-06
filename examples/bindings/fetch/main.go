package main

import "github.com/cookiengineer/gooey/bindings/fetch"
import "fmt"
import "time"

func main() {

	response, err := fetch.Fetch("http://localhost:3000/api/test", &fetch.Request{
		Method: fetch.MethodGet,
		Mode:   fetch.ModeSameOrigin,
	})

	if err == nil {
		fmt.Println(response)
	}

	for true {

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
