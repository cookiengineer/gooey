//go:build wasm

package storages

var global_localstorage *Storage

func init() {

	global_localstorage = &Storage{
		name: "localStorage",
	}

}

func GetLocalStorage() *Storage {
	return global_localstorage
}

