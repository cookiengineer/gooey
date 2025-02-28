package app

import "encoding/json"
import "errors"
import "gooey/storages"

type Storage struct {
}

func NewStorage() Storage {

	var storage Storage

	return storage

}

func (storage *Storage) Read(name string, schema any) error {

	var result error = nil

	buffer := storages.LocalStorage.GetItemBytes(name)

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
	storages.LocalStorage.RemoveItem(name)
}

func (storage *Storage) Write(name string, value any) error {

	var result error = nil

	buffer, err := json.Marshal(value)

	if err == nil {
		storages.LocalStorage.SetItem(name, buffer)
	} else {
		result = err
	}

	return result

}
