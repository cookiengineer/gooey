
# Common Errata and Limitations


## Promise APIs and Go Channels

Sometimes, the main thread will hang when using go channels inside a for loop.

The effect is visible that the line before will be executed, but then when trying to read a
channel, the program will never finish reading it. So every binding that is using a Promise
API behind the scenes can be affected by this scheduler bug.

The cause is likely the scheduler timing precision differences in Web Browsers and the
expectations on the internal go `syscall/js` / `wasm_exec.js` side, where channels seem
to not be properly deconstructed in some cases.

If you experience this bug, it's usually solved by wrapping the Promise API calls inside a
containing go routine so that the main thread doesn't hang up.

Example Code that will hang up:

```go
// DONT USE
for d := 0; d < len(dataset); d++ {

	payload, err1 := json.MarshalIndent(dataset[d], nil, "\t")

	fmt.Println("This will be printed")

	// XXX: This will hang up the main thread, when the Promise API is called
	// and the internal fetch_state channel is read
	response, err2 := fetch.Fetch("/api/whatever", &fetch.Request{
		Method:      fetch.MethodPost,
		Mode:        fetch.ModeSameOrigin,
		Cache:       fetch.CacheDefault,
		Credentials: fetch.CredentialsOmit,
		Redirect:    fetch.RedirectError,
		Headers:     map[string]string{
			"Accept":         "application/json",
			"Content-Type":   "application/json",
			"Content-Length": strconv.Itoa(len(payload)),
		},
		Body: bytes.NewReader(payload),
	})

	fmt.Println("This will not be printed")

}
```

Example Code that will work:

```go
for d := 0; d < len(dataset); d++ {

	go func(entry struct.Whatever) {

		payload, err1 := json.MarshalIndent(dataset[d], nil, "\t")

		fmt.Println("This will be printed")

		response, err2 := fetch.Fetch("/api/whatever", &fetch.Request{
			// Same as above
		})

		fmt.Println("This will be printed")

	}(dataset[d])

}
```


## Generic Methods and Components Typecasting

In Go there's still no generic methods available due to the ongoing discussion about
whether or not runtime boxing or comptime expansion is going to be implemented, which
means that the Web Component graph in the [components](../components) package is built
up using [interfaces.Component](../interfaces/Component.go).

This implies that the consuming side of the application needs to manually typecast the
components in the graph back to their struct references.

```go
import "example/actions"
import "example/schemas"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/layout"

// Example Controller
type Example struct {
	Main   *app.Main         `json:"main"`
	Schema *schemas.Examples `json:"schema"`
	View   *app.View         `json:"view"`
}

func NewExample(main *app.Main) Example {

	var controller Example

	controller.Main   = main
	controller.Schema = &schema.Examples{}
	controller.View   = view

	// body > main > section[data-name="example"] > article
	article, ok1 := controller.View.Content[0].(*layout.Article)

	if ok1 == true {

		// body > main > section[data-name="example"] > article > table
		table, ok2 := article.Content[0].(*content.Table)

		if ok2 == true {

			table.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {
				// Now you can use the Table API
			}))

		}

	}

	return controller

}
```

