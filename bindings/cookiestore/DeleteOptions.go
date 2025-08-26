package cookiestore

type DeleteOptions struct {
	Name        string `json:"name"`
	Domain      string `json:"domian"`
	Path        string `json:"path"`
	Partitioned bool   `json:"partitioned"`
}

func (options DeleteOptions) MapToJS() map[string]any {

	mapped := make(map[string]any)
	mapped["name"] = options.Name
	mapped["partitioned"] = options.Partitioned

	if options.Domain != "" {
		mapped["domain"] = options.Domain
	}

	if options.Path != "" {
		mapped["path"] = options.Path
	} else {
		mapped["path"] = "/"
	}

	return mapped

}
