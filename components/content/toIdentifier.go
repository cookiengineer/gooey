package content

func toIdentifier(value string) string {

	filtered := ""
	last_was_dash := false

	for _, chunk := range value {

		if chunk == 45 || chunk == 46 || chunk == 47 || chunk == 95 {

			if last_was_dash == false {
				filtered += "-"
				last_was_dash = true
			}

		} else if chunk >= 48 && chunk <= 57 {

			// Digits
			filtered += string(chunk)
			last_was_dash = false

		} else if chunk >= 58 && chunk <= 64 {

			if last_was_dash == false {
				filtered += "-"
				last_was_dash = true
			}

		} else if chunk >= 65 && chunk <= 90 {

			// Uppercase letters
			filtered += string(chunk)
			last_was_dash = false

		} else if chunk >= 91 && chunk <= 96 {

			if last_was_dash == false {
				filtered += "-"
				last_was_dash = true
			}

		} else if chunk >= 97 && chunk <= 122 {

			// Lowercase letters
			filtered += string(chunk)
			last_was_dash = false

		} else if chunk >= 123 {

			if last_was_dash == false {
				filtered += "-"
				last_was_dash = true
			}

		}

	}

	if len(filtered) > 0 {

		first_character := filtered[0]

		if (first_character >= 65 && first_character <= 90) || (first_character >= 97 && first_character <= 122) {

			// must start with A-Z or a-z
			return filtered

		} else {

			return ""

		}

	}

	return ""

}
