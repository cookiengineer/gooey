//go:build wasm

package utils

import "github.com/cookiengineer/gooey/bindings/dom"
import "strings"

func MatchesQuery(element *dom.Element, query string) bool {

	if element != nil {

		result := false

		if strings.Contains(query, " > ") {
			result = false
		} else if strings.Contains(query, " ") {
			result = false
		} else {

			if strings.Contains(query, "[") && strings.HasSuffix(query, "]") {

				matches_tagname   := false
				matches_attribute := false

				query_tagname := strings.ToLower(strings.TrimSpace(query[0:strings.Index(query, "[")]))

				if query_tagname == strings.ToLower(element.TagName) {
					matches_tagname = true
				}

				query_attribute := strings.TrimSpace(query[strings.Index(query, "[")+1:len(query)-1])

				if strings.Contains(query_attribute, "^=") {

					attribute := query_attribute[0:strings.Index(query_attribute, "^=")]
					value     := query_attribute[strings.Index(query_attribute, "^=")+2:]

					if strings.HasPrefix(element.GetAttribute(attribute), value) {
						matches_attribute = true
					}

				} else if strings.Contains(query_attribute, "$=") {

					attribute := query_attribute[0:strings.Index(query_attribute, "$=")]
					value     := query_attribute[strings.Index(query_attribute, "$=")+2:]

					if strings.HasSuffix(element.GetAttribute(attribute), value) {
						matches_attribute = true
					}

				} else if strings.Contains(query_attribute, "*=") {

					attribute := query_attribute[0:strings.Index(query_attribute, "*=")]
					value     := query_attribute[strings.Index(query_attribute, "*=")+2:]

					if strings.Contains(element.GetAttribute(attribute), value) {
						matches_attribute = true
					}

				} else if strings.Contains(query_attribute, "|=") {

					attribute := query_attribute[0:strings.Index(query_attribute, "^=")]
					value     := query_attribute[strings.Index(query_attribute, "^=")+2:]

					if element.GetAttribute(attribute) == value {
						matches_attribute = true
					} else if strings.HasPrefix(element.GetAttribute(attribute), value + "-") {
						matches_attribute = true
					}

				} else if strings.Contains(query_attribute, "=") {

					attribute := query_attribute[0:strings.Index(query_attribute, "=")]
					value     := query_attribute[strings.Index(query_attribute, "=")+1:]

					if element.GetAttribute(attribute) == value {
						matches_attribute = true
					}

				} else {

					attribute := strings.TrimSpace(query_attribute)

					if element.GetAttribute(attribute) != "" {
						matches_attribute = true
					}

				}

				if matches_tagname && matches_attribute {
					result = true
				}

			} else {

				if strings.ToLower(element.TagName) == strings.ToLower(query) {
					result = true
				}

			}

		}

		return result

	}

	return false

}
