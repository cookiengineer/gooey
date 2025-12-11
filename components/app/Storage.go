//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/storages"
import "encoding/json"
import "errors"

type Storage struct {
	storage *storages.Storage `json:"-"`
}

func NewStorage() *Storage {

	var storage Storage

	storage.storage = storages.GetLocalStorage()

	return &storage

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
