//go:build wasm

package dom

import "errors"
import "strconv"

func validateXMLName(name string) error {

	var err error = nil

	if len(name) > 0 {

		first := name[0]

		if first == ':' ||
			(first >= 'A' && first <= 'Z') ||
			first == '_' ||
			(first >= 'a' && first <= 'z') ||
			(first >= 0xC0 && first <= 0xD6) ||
			(first >= 0xD8 && first <= 0xF6) ||
			(first >= 0xF8 && first <= 0xFF) {

			valid := true

			var n int = 1

			for n = 1; n < len(name); n++ {

				chr := name[n]

				if first == '-' ||
					first == '.' ||
					(first >= '0' && first <= '9') ||
					chr == ':' ||
					(first >= 'A' && first <= 'Z') ||
					first == '_' ||
					(first >= 'a' && first <= 'z') ||
					first == '·' ||
					(first >= 0xC0 && first <= 0xD6) ||
					(first >= 0xD8 && first <= 0xF6) ||
					(first >= 0xF8 && first <= 0xFF) {
					// Do Nothing
				} else {
					valid = false
					break
				}

			}

			if valid == false {
				err = errors.New("InvalidCharacterError: character " + strconv.Itoa(n) + " is invalid")
			}

		} else {
			err = errors.New("InvalidCharacterError: character 0 invalid")
		}

	}

	return err

}
