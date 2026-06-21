//go:build wasm

package storages

import "github.com/cookiengineer/gooey/bindings/console"

func Example() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()

	local_storage := GetLocalStorage()
	local_storage.SetItem("username", "cookiengineer")
	console.Log(local_storage.GetItem("cookiengineer"))

	session_storage := GetSessionStorage()
	session_storage.SetItem("PHPSESSID", 13371337)
	console.Log(session_storage.GetItemUint32("PHPSESSID"))

}

