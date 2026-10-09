package cookie

type DeleteOptions struct {
	Name        string  `json:"name"`
	Domain      *string `json:"domian"`
	Path        *string `json:"path"`
	Partitioned bool    `json:"partitioned"`
}

func (deleteOptions DeleteOptions) MapToJS() map[string]any {
	mapped := make(map[string]any)

	mapped["name"] = deleteOptions.Name

	if deleteOptions.Domain != nil {
		mapped["domain"] = *deleteOptions.Domain
	}

	path := "/"
	if deleteOptions.Path != nil {
		path = *deleteOptions.Path
	}
	mapped["path"] = path

	mapped["partitioned"] = deleteOptions.Partitioned

	return mapped
}
