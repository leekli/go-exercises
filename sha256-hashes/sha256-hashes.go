// SHA256 Hashes
// - SHA256 hashes are frequently used to compute short identities for binary or text blobs. For example, TLS/SSL certificates use SHA256 to compute a certificate’s signature.
// - `crypto/sha256` package (other hash functions available in crypto/*)

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	shaStr := "sha256 this string"

	// New sha256 hash function
	newSha := sha256.New()

	// Need co-erce string to bytes
	newSha.Write([]byte(shaStr))

	bs := newSha.Sum(nil)

	fmt.Println("Original string:", shaStr)
    fmt.Printf("SHA256 Hash: %x\n", bs)
}