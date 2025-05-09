package aescbc

import "syscall/js"

type CryptoKey struct {
	Type        CryptoKeyType `json:"type"`
	Extractable bool          `json:"extractable"`
	Usages      []string      `json:"usages"`
	Value       *js.Value     `json:"value"`
	Algorithm   struct {
		Name   string `json:"name"`
		Length int    `json:"length"`
	} `json:"algorithm"`
}

func ToCryptoKey(value js.Value) *CryptoKey {

	var key CryptoKey

	usages := make([]string, 0)
	usages_array := value.Get("usages")

	for u := 0; u < usages_array.Length(); u++ {
		usages = append(usages, usages_array.Index(u).String())
	}

	key.Algorithm.Name = value.Get("algorithm").Get("name").String()
	key.Algorithm.Length = value.Get("algorithm").Get("length").Int()
	key.Type = CryptoKeyType(value.Get("type").String())
	key.Extractable = value.Get("extractable").Bool()
	key.Usages = usages
	key.Value = &value

	return &key

}
