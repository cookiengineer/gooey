//go:build wasm

package webgl

import "syscall/js"

type Program struct {
	Linked  bool      `json:"linked"`
	InfoLog string    `json:"infoLog"`
	Value   *js.Value `json:"value"`
}

func ToProgram(value js.Value) *Program {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var program Program

	program.Linked = false
	program.InfoLog = ""
	program.Value = &value

	return &program

}
