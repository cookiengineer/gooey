//go:build wasm

package storages

var LocalStorage Storage

func init() {

	LocalStorage = Storage{
		name: "localStorage",
	}

}
