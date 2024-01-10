// Package algorithms provides various cryptographic algorithm implementations.
// This file includes the implementation for the Keccak-256 hashing algorithm.
package algorithms

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

// Keccak256Hash computes the Keccak-256 hash of the input string.
// Keccak-256 is a member of the SHA-3 family of cryptographic hash functions and is used in the Ethereum blockchain.
//
// @param input The input string to hash.
// @return The hexadecimal representation of the Keccak-256 hash of the input.
func Keccak256Hash(input string) string {
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
