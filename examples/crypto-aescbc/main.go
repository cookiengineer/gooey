package main

import "gooey"
import "gooey/crypto"
import "gooey/crypto/aescbc"
import "encoding/hex"
import "encoding/json"
import "fmt"
import "strings"
import "time"

func main() {

	key, err1 := aescbc.GenerateKey(256, true, []string{
		"encrypt",
		"decrypt",
	})

	if err1 != nil {
		fmt.Println(err1)
	}

	iv      := crypto.GetRandomValues(16)
	message := []byte("This was the plaintext message as a string")

	details1, _ := json.MarshalIndent(key, "", "\t")
	element1 := gooey.Document.QuerySelector("#generated-key")
	element1.SetInnerHTML(string(details1))

	encrypted, err2 := aescbc.Encrypt(iv, key, message)

	if err2 != nil {
		fmt.Println(err2)
	}

	details2 := make([]string, 0)

	for e := 0; e < len(encrypted); e++ {
		details2 = append(details2, "0x" + hex.EncodeToString([]byte{encrypted[e]}))
	}

	fmt.Println(encrypted, err2)

	element2 := gooey.Document.QuerySelector("#encrypted-buffer")
	element2.SetInnerHTML(strings.Join(details2, " "))

	decrypted, err3 := aescbc.Decrypt(iv, key, encrypted)

	if err3 != nil {
		fmt.Println(err3)
	}

	details3 := string(decrypted)
	element3 := gooey.Document.QuerySelector("#decrypted-buffer")
	element3.SetInnerHTML(details3)


	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
