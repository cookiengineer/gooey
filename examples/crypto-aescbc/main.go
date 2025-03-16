package main

import gooey "github.com/cookiengineer/gooey/pkg"
import "github.com/cookiengineer/gooey/pkg/console"
import "github.com/cookiengineer/gooey/pkg/crypto"
import "github.com/cookiengineer/gooey/pkg/crypto/aescbc"
import "encoding/hex"
import "encoding/json"
import "strings"
import "time"

func main() {

	// This is the example message
	message := []byte("This was the plaintext message as a string")

	// Step 1: Generate CryptoKey
	key, err1 := aescbc.GenerateKey(256, true, []string{
		"encrypt",
		"decrypt",
	})

	if err1 != nil {
		console.Error(err1)
	}

	details1, _ := json.MarshalIndent(key, "", "\t")
	element1 := gooey.Document.QuerySelector("#generated-cryptokey")
	element1.SetInnerHTML(string(details1))

	// Step 2: Generate IV (Initialization Vector)
	iv       := crypto.GetRandomValues(16)
	details2 := make([]string, 0)

	for i := 0; i < len(iv); i++ {
		details2 = append(details2, "0x" + hex.EncodeToString([]byte{iv[i]}))
	}

	element2 := gooey.Document.QuerySelector("#generated-iv")
	element2.SetInnerHTML(strings.Join(details2, " "))


	// Step 3: Encrypt with IV and CryptoKey
	encrypted, err3 := aescbc.Encrypt(iv, key, message)
	details3        := make([]string, 0)

	if err3 != nil {
		console.Error(err3)
	}

	for e := 0; e < len(encrypted); e++ {
		details3 = append(details3, "0x" + hex.EncodeToString([]byte{encrypted[e]}))
	}

	element3 := gooey.Document.QuerySelector("#encrypted-buffer")
	element3.SetInnerHTML(strings.Join(details3, " "))


	// Step 4: Decrypt with IV and CryptoKey
	decrypted, err4 := aescbc.Decrypt(iv, key, encrypted)

	if err4 != nil {
		console.Error(err4)
	}

	details4 := string(decrypted)
	element4 := gooey.Document.QuerySelector("#decrypted-buffer")
	element4.SetInnerHTML(details4)


	// Step 5: Export Key
	exported, err5 := aescbc.ExportKey("jwk", *key)

	if err5 != nil {
		console.Error(err5)
	}

	details5 := string(exported)
	element5 := gooey.Document.QuerySelector("#exported-keydata")
	element5.SetInnerHTML(details5)


	// Step 6: Import Key
	imported, err6 := aescbc.ImportKey("jwk", exported, true, []string{
		"encrypt",
		"decrypt",
	})

	if err6 != nil {
		console.Error(err6)
	}

	details6, _ := json.MarshalIndent(imported, "", "\t")
	element6 := gooey.Document.QuerySelector("#imported-cryptokey")
	element6.SetInnerHTML(string(details6))


	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
