
# Promise APIs and Go Channels

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
