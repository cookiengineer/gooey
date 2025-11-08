package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/cookiestore"
import "github.com/cookiengineer/gooey/bindings/timers"
import "time"

func main() {

	console := console.GetConsole()

	cookiestore.Set(cookiestore.SetOptions{
		Name:  "cookie-monster",
		Value: "Me want cookies!",
	})

	timers.SetTimeout(func() {

		cookie, err := cookiestore.Get(cookiestore.GetOptions{
			Name: "cookie-monster",
		})

		if err == nil {
			console.Log(cookie)
		} else {
			console.Error(err)
		}

	}, 1000)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
