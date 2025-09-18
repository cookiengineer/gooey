
# Common Errata and Limitations


## Deadlocks

Sometimes, the main thread will hang when using `wasm` as a compilation target.

Go's scheduler is synchronous, meaning that if you have a function that executes something,
it will block the main thread by default. In order to run things in parallel,
go has the `go func(){}()` or go routine concept which is non-blocking.

The deadlock effect is visible that the line before it will be executed, but then when
trying to read a channel or executing a `waitGroup.Wait()`, the program will never finish
and the CPU load will spike to an endless for loop.

Usually the Browser Tab will be non-reactive as well, because the scheduler's main loop
will just endlessly check for something that will never evaluate to true.


### Deadlocks with Promise APIs

Every Promise API in the Browser has to use a `chan` (or channel) to represent its state
of whether it was successful by executing the `then(func(){})` callback or unsuccessful by executing
the `catch(func(){})` callback.

Keep in mind that all implementations of `gooey/bindings` that use Promises in the Web
Browser will have a channel that could be blocking the main thread and result into a deadlock.

If you experience this bug, it's usually solved by wrapping the Promise API calls inside a
containing `go func(){}()` wrapper so that the main thread doesn't hang up in a deadlock.

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


### Deadlocks with WaitGroups

Things get a little complicated when you rely on something like a `sync.Waitgroup` in Go
because its `Wait()` function is blocking the main thread and will result in a deadlock
of the main loop that is provided by the `wasm_exec.js` function.

Example Code that will hang up:

```go
// DONT USE
waitgroup := sync.WaitGroup{}

for i := 0; i < 10; i++ {
	waitgroup.Add(1)

	go func(i int) {
		time.Sleep(i * time.Second)
		fmt.Println("Waited " + strconv.Itoa(i) + " seconds")
		defer waitgroup.Done()
	}(i)

}

// waitgroup.Wait() will deadlock the main thread
waitgroup.Wait()
fmt.Println("Finished")
```

Example Code that will work:

```go
go func() {

	waitgroup := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		waitgroup.Add(1)

		go func(i int) {
			time.Sleep(i * time.Second)
			fmt.Println("Waited " + strconv.Itoa(i) + " seconds")
			defer waitgroup.Done()
		}(i)

	}

	waitgroup.Wait()
	fmt.Println("Finished")

}()
```


### Deadlocks with Channels

If you read from a channel in Go, it will cause a deadlock in the `wasm_exec.js` scheduler.
This means that you have to wrap all channel-using code into a `go func(){}()` wrapper.

This is related to Promise APIs, too, because they rely on using channels to be able to both
evaluate the result or error state from `then()` and `catch()` callbacks.

Example Code that will hang up:

```go
// DONT USE
func main() {

	channel := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		channel <- true
	}()

	result := <-done

	fmt.Println("Deadlock of main thread happens before this")
	fmt.Println(result)

}
```

Example code that will work:

```go
func main() {

	go func() {
		channel := make(chan bool)

		go func() {
			time.Sleep(5 * time.Second)
			channel <- true
		}()

		result := <-done

		fmt.Println("No deadlock of main thread")
		fmt.Println(result)
	}()

}
```


### Debugging Deadlocks

Mistakes happen, and it's a little painful to debug deadlocks because the Go scheduler
in `wasm_exec.js` won't complete and therefore you don't know where exactly your code
will deadlock.

However, you can get better debug information by using the `runtime.Stack()` method and
by wrapping a little helper method to do it. This will print out all running go routines
and show the stacktraces of each of them.

```go
func DebugDeadlockAfter(seconds int) {
	time.Sleep(seconds * time.Second)
	buffer := make([]byte, 1<<16)
	runtime.Stack(buf, true)

	fmt.Println(string(buf))
}

func main() {

	// Usage example
	go DebugDeadlockAfter(10)

	waitgroup := sync.WaitGroup{}
	waitgroup.Add(1)
	go func() {
		time.Sleep(30 * time.Second)
		defer waitgroup.Done()
	}()

	// Deadlock main thread on purpose
	waitgroup.Wait()

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


