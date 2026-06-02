//go:build wasm

package xhr

import "github.com/cookiengineer/gooey/bindings/console"
import "strings"
import "time"

func Example() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "strconv"
	// import "strings"
	// import "time"

	console := console.GetConsole()
	xhr := NewXMLHttpRequest()

	xhr.Timeout = 10 * time.Second
	xhr.WithCredentials = false

	xhr.OnLoad = func(status int, response []byte) {

		content_type := xhr.GetResponseHeader("Content-Type")

		if content_type == "application/json" {
			console.Info(status, string(response))
		} else {
			console.Warn(status, response)
		}

	}

	xhr.OnError = func() {
		console.Error("XMLHttpRequest network error!")
	}

	xhr.OnTimeout = func() {
		console.Warn("XMLHttpRequest timed out!")
	}

	xhr.Open(xhr.MethodPost, "https://example.com/api")

	body := strings.NewReader("{\"message\": \"Hello, world!\"}")

	xhr.SetRequestHeader("Accept", "application/json")
	xhr.SetRequestHeader("Content-Type", "application/json")
	xhr.SetRequestHeader("Content-Length", strconv.Itoa(28))
	xhr.Send(body)

}
