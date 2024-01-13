// Base64 Encoding
// - part of `encoding/decoding` package

package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

    data := "abc123!?$*&()'-=@~"

	// Encoding string `data` to base64
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

	// Decode the string back
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}