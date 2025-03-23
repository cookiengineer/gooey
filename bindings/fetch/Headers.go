//go:build wasm

package fetch

import "io"
import "net/textproto"
import "strings"

type Headers map[string][]string

func (headers Headers) Append(key string, value string) {
	textproto.MIMEHeader(headers).Add(key, value)
}

func (headers Headers) Delete(key string) {
	textproto.MIMEHeader(headers).Del(key)
}

func (headers Headers) Get(key string) string {
	return textproto.MIMEHeader(headers).Get(key)
}

func (headers Headers) Has(key string) bool {

	var result bool = false

	_, ok := headers[key]

	if ok == true {
		result = true
	}

	return result

}

func (headers Headers) Set(key string, value string) {
	textproto.MIMEHeader(headers).Set(key, value)
}

func (headers Headers) Write(writer io.Writer) error {

	string_writer, ok := writer.(io.StringWriter)

	if ok == false {
		string_writer = headersStringWriter{writer}
	}

	keyvalues := sortHeaderKeyValues(headers)

	for _, keyvaluepair := range keyvalues {

		for _, value := range keyvaluepair.values {

			value = strings.ReplaceAll(value, "\n", " ")
			value = strings.ReplaceAll(value, "\r", " ")
			value = textproto.TrimString(value)

			_, err := string_writer.WriteString(keyvaluepair.key + ": " + value + "\r\n")

			if err != nil {
				return err
			}

		}

	}

	return nil

}
