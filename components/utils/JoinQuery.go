//go:build wasm

package utils

import "strings"

func JoinQuery(selectors []string) string {
	return strings.Join(selectors, " > ")
}
