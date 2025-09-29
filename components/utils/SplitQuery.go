//go:build wasm

package utils

import "strings"

func SplitQuery(query string) []string {

	result := make([]string, 0)
	tmp1 := strings.TrimSpace(query)

	if strings.Contains(tmp1, " > ") {

		tmp2 := strings.Split(tmp1, " > ")

		if len(tmp2) > 0 {

			for _, chunk := range tmp2 {

				tmp := strings.TrimSpace(chunk)

				if tmp != "" {
					result = append(result, tmp)
				}

			}

		}

	} else {

		if tmp1 != "" {
			result = append(result, tmp1)
		}

	}

	return result

}
