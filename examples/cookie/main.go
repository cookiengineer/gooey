//go:build wasm

package main

import (
	"fmt"
	"gooey/cookie"
	"time"
)

func main() {
	cookieStore := cookie.CookieStore
	cookies, err := cookieStore.GetAll(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(cookies)
	for true {

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
