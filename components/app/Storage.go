//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/storages"
import "github.com/cookiengineer/gooey/components/reactive"
import "encoding/json"
import "errors"

type Storage struct {
	storage *storages.Storage `json:"-"`
	store   *reactive.Store   `json:"-"`
}

func NewStorage() *Storage {

	var storage Storage

	storage.storage = storages.GetLocalStorage()
	storage.store   = reactive.NewStore()

	return &storage

}

// Get returns the in-memory value and whether it has been set.
func (storage *Storage) Get(name string) (any, bool) {
	return storage.store.Get(name)
}

// Revision returns the current revision for a name.
func (storage *Storage) Revision(name string) uint64 {
	return storage.store.Revision(name)
}

// Subscribe registers a watcher for a name and returns an unsubscribe func.
func (storage *Storage) Subscribe(name string, callback func(value any)) func() {
	return storage.store.Subscribe(name, callback)
}

// Update sets the in-memory value, bumps the revision, persists it to
// LocalStorage and then notifies all watchers.
func (storage *Storage) Update(name string, value any) error {

	storage.store.Update(name, value)

	err := storage.Write(name, value)

	storage.store.Notify(name)

	return err

}

func (storage *Storage) Read(name string, schema any) error {

	var result error = nil

	buffer := storage.storage.GetItemBytes(name)

	if len(buffer) > 0 {

		err := json.Unmarshal(buffer, &schema)

		if err != nil {
			result = err
		}

	} else {
		result = errors.New("\"" + name + "\" does not exist in LocalStorage")
	}

	return result

}

func (storage *Storage) Remove(name string) {
	storage.storage.RemoveItem(name)
}

func (storage *Storage) Write(name string, value any) error {

	var result error = nil

	buffer, err := json.Marshal(value)

	if err == nil {
		storage.storage.SetItem(name, buffer)
	} else {
		result = err
	}

	return result

}
