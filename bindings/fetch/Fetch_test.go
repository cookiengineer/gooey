//go:build wasm

package fetch

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleFetch() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	response, err := Fetch("/api/example", &Request{
		Method:  MethodGet,
		Headers: map[string]string{
			"Accept": "application/json",
			"X-Application": "gooey-example",
		},
		Mode:           ModeCORS,
		Credentials:    CredentialsOmit,
		Cache:          CacheDefault,
		Redirect:       RedirectError,
		Referrer:       "http://localhost:1337",
		ReferrerPolicy: ReferrerPolicyStrictOriginWhenCrossOrigin,
		// Integrity:      "sha512-...", // Integrity Hash for CDN requests
		KeepAlive:      false,
	})

	if err != nil {
		console.Log(response)
	}

}
