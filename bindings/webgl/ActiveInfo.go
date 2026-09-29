//go:build wasm

package webgl

import "syscall/js"

type ActiveInfo struct {
	Name  string    `json:"name"`
	Size  int       `json:"size"`
	Type  DataType  `json:"type"`
	Value *js.Value `json:"value"`
}

func ToActiveInfo(value js.Value) *ActiveInfo {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var info ActiveInfo

	info.Name = value.Get("name").String()
	info.Size = value.Get("size").Int()
	info.Type = DataType(value.Get("type").Int())
	info.Value = &value

	return &info

}
