//go:build wasm

package webgl

import "syscall/js"

type Shader struct {
	Type     ShaderType `json:"type"`
	Source   string     `json:"source"`
	Compiled bool       `json:"compiled"`
	InfoLog  string     `json:"infoLog"`
	Value    *js.Value  `json:"value"`
}

func ToShader(value js.Value) *Shader {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var shader Shader

	shader.Type = ShaderTypeVertex
	shader.Source = ""
	shader.Compiled = false
	shader.InfoLog = ""
	shader.Value = &value

	return &shader

}
