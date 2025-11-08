//go:build wasm

package location

func onchange(location *Location) {

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

