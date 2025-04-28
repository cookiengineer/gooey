//go:build wasm

package cookiestore

import "syscall/js"

type Cookie struct {
	Domain      string   `json:"domain"`
	Expires     int      `json:"expires"`
	Name        string   `json:"name"`
	Partitioned bool     `json:"partitioned"`
	Path        string   `json:"path"`
	SameSite    SameSite `json:"sameSite"`
	Secure      bool     `json:"secure"`
	Value       string   `json:"value"`
}

func ToCookie(value js.Value) Cookie {

	var cookie Cookie

	domain := value.Get("domain")
	expires := value.Get("expires")
	partitioned := value.Get("partitioned")
	path := value.Get("path")
	samesite := value.Get("sameSite")
	secure := value.Get("secure")

	if !domain.IsNull() && !domain.IsUndefined() {
		cookie.Domain = domain.String()
	}

	if !expires.IsNull() && !expires.IsUndefined() {
		// milliseconds are stripped off
		cookie.Expires = int64(expires.Float())
	}

	if !partitioned.IsNull() && !partitioned.IsUndefined() {
		cookie.Partitioned = partitioned.Bool()
	}

	if !path.IsNull() && !path.IsUndefined() {
		cookie.Path = path.String()
	}

	if !samesite.IsNull() && !samesite.IsUndefined() {
		cookie.SameSite = SameSite(samesite.String())
	}

	if !secure.IsNull() && !secure.IsUndefined() {
		cookie.Secure = secure.Bool()
	}

	cookie.Name = value.Get("name").String()
	cookie.Value = value.Get("value").String()

	// XXX: Currently unstable properties, see above
	// cookie.Domain      = value.Get("domain").String()
	// cookie.Expires     = value.Get("expires").Int()
	// cookie.Partitioned = value.Get("partitioned").Bool()
	// cookie.Path        = value.Get("path").String()
	// cookie.Secure      = value.Get("secure").Bool()

	return cookie

}
