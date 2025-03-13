//go:build wasm

package fetch

import "io"

type headersStringWriter struct {
	w io.Writer
}

func (writer headersStringWriter) WriteString(value string) (n int, err error) {
	return writer.w.Write([]byte(value))
}
