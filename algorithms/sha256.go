// Package algorithms provides various cryptographic algorithm implementations.
// This file includes implementations for SHA-2 and SHA-3 hashing algorithms.
package algorithms

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// Sha256Hash computes the SHA-256 hash of the input string.
// SHA-256 is part of the SHA-2 family of cryptographic hash functions designed by the NSA.
// It produces a 256-bit (32-byte) hash value, typically rendered as a hexadecimal number, 64 digits long.
//
// @param input The input string to hash.
// @return The hexadecimal representation of the SHA-256 hash of the input.
func Sha256Hash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}

// Sha3Hash computes the SHA3-256 hash of the input string.
// SHA-3 is a subset of the broader cryptographic primitive family Keccak, designed by Guido Bertoni, Joan Daemen, Michaël Peeters, and Gilles Van Assche.
// SHA3-256 also produces a 256-bit (32-byte) hash value.
//
// @param input The input string to hash.
// @return The hexadecimal representation of the SHA3-256 hash of the input.
func Sha3Hash(input string) string {
	hash := sha3.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}
