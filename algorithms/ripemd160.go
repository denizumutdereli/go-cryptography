// Package algorithms provides various cryptographic algorithm implementations.
// This file includes the implementation for the RIPEMD-160 hashing algorithm.
package algorithms

import (
	"crypto/sha256"
	"fmt"

	// old fashion
	"golang.org/x/crypto/ripemd160"
)

// GenerateRIPEMD160Hash generates a RIPEMD-160 hash from an input string.
// Typically, this is used in Bitcoin for hashing the public key to create addresses.
func GenerateRIPEMD160Hash(input string) string {
	// First, hash the input using SHA-256
	sha256Hash := sha256.Sum256([]byte(input))

	// Then, hash the SHA-256 hash using RIPEMD-160
	ripemd160Hasher := ripemd160.New()
	_, err := ripemd160Hasher.Write(sha256Hash[:])
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", ripemd160Hasher.Sum(nil))
}
