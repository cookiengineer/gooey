package cookiestore

type GetOptions struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (options GetOptions) MapToJS() map[string]any {

	mapped := make(map[string]any)
	mapped["name"] = options.Name
	mapped["url"] = options.Url

	return mapped

}
