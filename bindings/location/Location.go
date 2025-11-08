//go:build wasm

package location

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

func (location *Location) Assign(url string) {

	if location.Value != nil && !location.Value.IsNull() && !location.Value.IsUndefined() {
		location.Value.Call("assign", js.ValueOf(url))
		onchange(location)
	}

}

func (location *Location) Reload() {

	if location.Value != nil && !location.Value.IsNull() && !location.Value.IsUndefined() {
		location.Value.Call("reload")
	}

}

func (location *Location) Replace(url string) {

	if location.Value != nil && !location.Value.IsNull() && !location.Value.IsUndefined() {
		location.Value.Call("replace", js.ValueOf(url))
		onchange(location)
	}

}
