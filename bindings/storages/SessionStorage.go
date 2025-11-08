//go:build wasm

package storages

var global_sessionstorage *Storage

func init() {

	global_sessionstorage = &Storage{
		name: "sessionStorage",
	}

}

func GetSessionStorage() *Storage {
	return global_sessionstorage
}

