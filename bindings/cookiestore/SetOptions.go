package cookiestore

type SetOptions struct {
	Domain      *string   `json:"domain"`
	Expires     *int      `json:"expires"`
	Name        string    `json:"name"`
	Partitioned bool      `json:"partitioned"`
	Path        *string   `json:"path"`
	SameSite    *SameSite `json:"sameSite"`
	Value       string    `json:"value"`
}

func (options SetOptions) MapToJS() map[string]any {

	mapped := make(map[string]any)
	mapped["name"] = options.Name
	mapped["partitioned"] = options.Partitioned
	mapped["value"] = options.Value

	if options.Domain != nil {
		mapped["domain"] = *options.Domain
	}

	if options.Expires != nil {
		mapped["expires"] = options.Expires
	} else {
		mapped["expires"] = nil
	}

	if options.Path != nil {
		mapped["path"] = *options.Path
	} else {
		mapped["path"] = "/"
	}

	if options.SameSite != nil {
		mapped["samesite"] = options.SameSite.String()
	} else {
		mapped["samesite"] = SameSiteStrict.String()
	}

	return mapped

}
