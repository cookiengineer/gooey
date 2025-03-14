//go:build wasm

package history

type HistoryState struct {
	State *map[string]any `json:"state"`
	Title string          `json:"title"`
	URL   string          `json:"url"`
	value *js.Value       `json:"value"`
}

