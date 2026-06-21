//go:build wasm

package location

import "github.com/cookiengineer/gooey/bindings/quirks"
import "errors"
import "syscall/js"

var global_location *Location

func init() {
	global_location = GetLocation()
}

type Location struct {
	Href     string    `json:"href"`
	Protocol string    `json:"protocol"`
	Host     string    `json:"host"`
	Hostname string    `json:"hostname"`
	Port     string    `json:"port"`
	Pathname string    `json:"pathname"`
	Search   string    `json:"search"`
	Hash     string    `json:"hash"`
	Origin   *string   `json:"origin"`
	Value    *js.Value `json:"value"`
}

// Returns the global Location instance.
func GetLocation() *Location {

	if global_location != nil {

		return global_location

	} else {

		value := js.Global().Get("location")
		check := value.Get("origin").String()

		var origin *string

		if check != "null" {
			origin = &check
		} else {
			origin = nil
		}

		location := Location{
			Href:     value.Get("href").String(),
			Protocol: value.Get("protocol").String(),
			Host:     value.Get("host").String(),
			Hostname: value.Get("hostname").String(),
			Port:     value.Get("port").String(),
			Pathname: value.Get("pathname").String(),
			Search:   value.Get("search").String(),
			Hash:     value.Get("hash").String(),
			Origin:   origin,
			Value:    &value,
		}

		return &location

	}

}

// Assigns a new URL to the window and forces a (re-)load while preserving a Browser History entry.
func (location *Location) Assign(url string) error {

	if location.Value != nil {

		if location.Value.IsNull() == false && location.Value.IsUndefined() == false {

			err := quirks.GoTryCatch(func() {
				location.Value.Call("assign", js.ValueOf(url))
			})

			if err == nil {

				onchange(location)
				return nil

			} else {
				return err
			}

		} else {
			return errors.New("Error: Location API not supported")
		}

	} else {
		return errors.New("Error: Location API not supported")
	}

}

// Reloads the current URL in the current window.
func (location *Location) Reload() error {

	if location.Value != nil {

		if location.Value.IsNull() == false && location.Value.IsUndefined() == false {

			err := quirks.GoTryCatch(func() {
				location.Value.Call("reload")
			})

			if err == nil {
				return nil
			} else {
				return err
			}

		} else {
			return errors.New("Error: Location API not supported")
		}

	} else {
		return errors.New("Error: Location API not supported")
	}

}

// Replaces the current URL in the current window without preserving a Browser History entry.
func (location *Location) Replace(url string) error {

	if location.Value != nil {

		if location.Value.IsNull() == false && location.Value.IsUndefined() == false {

			err := quirks.GoTryCatch(func() {
				location.Value.Call("replace", js.ValueOf(url))
			})

			if err == nil {

				onchange(location)
				return nil

			} else {
				return err
			}

		} else {
			return errors.New("Error: Location API not supported")
		}

	} else {
		return errors.New("Error: Location API not supported")
	}

}
