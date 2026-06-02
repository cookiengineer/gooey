//go:build wasm

package cookiestore

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleCookieStore_Delete() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	cookiestore := GetCookieStore()

	err := cookiestore.Delete(DeleteOptions{
		Name:        "username",
		Domain:      "example.com",
		Path:        "/login",
		Partitioned: true,
	})

	if err == nil {
		console.Info("Cookie \"username\" deleted!")
	} else {
		console.Error(err)
	}

}

func ExampleCookieStore_Get() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	cookiestore := GetCookieStore()

	cookie, err := cookiestore.Get(GetOptions{
		Name: "username",
		Url:  "https://example.com/login",
	})

	if err == nil {

		// Cookie{
		//   Domain:      "example.com",
		//   Expires:     1767139200000,
		//   Name:        "username",
		//   Partitioned: true,
		//   Path:        "/login",
		//   SameSite:    SameSite("strict"),
		//   Secure:      true,
		//   Value:       "cookiengineer",
		// }

		console.Info(cookie)

	} else {
		console.Error(err)
	}

}

func ExampleCookieStore_GetAll() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	cookiestore := GetCookieStore()

	cookies, err := cookiestore.GetAll(GetOptions{
		Name: "username",
		Url:  "https://example.com/login",
	})

	if err == nil {

		// []Cookie{
		//   {
		//     Domain:      "example.com",
		//     Expires:     1767139200000,
		//     Name:        "username",
		//     Partitioned: true,
		//     Path:        "/login",
		//     SameSite:    SameSite("strict"),
		//     Secure:      true,
		//     Value:       "cookiengineer",
		//   }, {
		//     Domain:      "example.com",
		//     Expires:     1767139200000,
		//     Name:        "PHPSESSID",
		//     Partitioned: true,
		//     Path:        "/login",
		//     SameSite:    SameSite("strict"),
		//     Secure:      false,
		//     Value:       "13371337",
		//   }
		// }

		console.Info(cookies)

	} else {
		console.Error(err)
	}

}

func ExampleCookieStore_Set() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	cookiestore := GetCookieStore()

	err := cookiestore.Set(SetOptions{
		Domain:      "example.com",
		Expires:     1767139200000,
		Name:        "username",
		Partitioned: true,
		Path:        "/login",
		SameSite:    &SameSite("strict"),
		Secure:      true,
		Value:       "cookiengineer",
	})

	if err == nil {
		console.Info("Cookie \"username\" has been stored!")
	} else {
		console.Error(err)
	}

}
