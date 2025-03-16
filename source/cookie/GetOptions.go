package cookie

type GetOptions struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (getOptions GetOptions) MapToJS() map[string]any {
	mapped := make(map[string]any)
	mapped["name"] = getOptions.Name
	mapped["url"] = getOptions.Url
	return mapped
}
