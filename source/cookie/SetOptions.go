package cookie

type SetOptions struct {
	Domain      *string   `json:"domain"`
	Expires     *int      `json:"expires"`
	Name        string    `json:"name"`
	Partitioned bool      `json:"partitioned"`
	Path        *string   `json:"path"`
	SameSite    *SameSite `json:"sameSite"`
	Value       string    `json:"value"`
}

func (setOptions SetOptions) MapToJS() map[string]any {
	mapped := make(map[string]any)

	mapped["name"] = setOptions.Name
	mapped["value"] = setOptions.Value

	if setOptions.Domain != nil {
		mapped["domain"] = *setOptions.Domain
	}

	mapped["expires"] = nil
	if setOptions.Expires != nil {
		mapped["expires"] = setOptions.Expires
	}

	mapped["partitioned"] = setOptions.Partitioned

	path := "/"
	if setOptions.Path != nil {
		path = *setOptions.Path
	}
	mapped["path"] = path

	sameSite := Strict
	if setOptions.SameSite != nil {
		sameSite = *setOptions.SameSite
	}
	mapped["sameSite"] = sameSite.ToString()

	return mapped
}
