package cookie

type Cookie struct {
	Domain      string   `json:"domain"`
	Expires     int64    `json:"expires"`
	Name        string   `json:"name"`
	Partitioned bool     `json:"partitioned"`
	Path        string   `json:"path"`
	SameSite    SameSite `json:"sameSite"`
	Secure      bool     `json:"secure"`
	Value       string   `json:"value"`
}
