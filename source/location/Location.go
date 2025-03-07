package location

import "syscall/js"

var Location location

func onchange(location *location) {

	location.Href = location.Value.Get("href").String()
	location.Protocol = location.Value.Get("protocol").String()
	location.Host = location.Value.Get("host").String()
	location.Hostname = location.Value.Get("hostname").String()
	location.Port = location.Value.Get("port").String()
	location.Pathname = location.Value.Get("pathname").String()
	location.Search = location.Value.Get("search").String()
	location.Hash = location.Value.Get("hash").String()

	origin := location.Value.Get("origin").String()

	if origin != "null" {
		location.Origin = &origin
	} else {
		location.Origin = nil
	}

}

func init() {

	value := js.Global().Get("location")
	check := value.Get("origin").String()

	var origin *string

	if check != "null" {
		origin = &check
	} else {
		origin = nil
	}

	Location = location{
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

}

type location struct {
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

func (location *location) Assign(url string) {

	if location.Value != nil && !location.Value.IsNull() && !location.Value.IsUndefined() {
		location.Value.Call("assign", js.ValueOf(url))
		onchange(location)
	}

}

func (location *location) Reload() {

	if location.Value != nil && !location.Value.IsNull() && !location.Value.IsUndefined() {
		location.Value.Call("reload")
	}

}

func (location *location) Replace(url string) {

	if location.Value != nil && !location.Value.IsNull() && !location.Value.IsUndefined() {
		location.Value.Call("replace", js.ValueOf(url))
		onchange(location)
	}

}
