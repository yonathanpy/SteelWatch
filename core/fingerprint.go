package core

import (
	"crypto/sha1"
	"fmt"
)

// Fingerprint generates a unique identifier for a client request
func Fingerprint(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return fmt.Sprintf("%x", h.Sum(nil))
}
