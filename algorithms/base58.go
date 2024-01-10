// Package algorithms provides various cryptographic algorithm implementations.
// This file includes the implementation for the Base58 encoding algorithm.
package algorithms

import (
	"bytes"
	"math/big"
)

const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Base58Encode encodes a byte array into a Base58 string.
func Base58Encode(input []byte) string {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}

	// Add leading zero bytes
	for _, b := range input {
		if b == 0x00 {
			result = append(result, base58Alphabet[0])
		} else {
			break
		}
	}

	// Reverse bytes
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Base58Decode decodes a Base58 string into a byte array.
func Base58Decode(input string) []byte {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte([]byte(base58Alphabet), byte(b))
		if charIndex == -1 {
			return []byte("")
		}
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()

	if len(input) > 0 {
		for _, b := range input {
			if rune(b) != rune(base58Alphabet[0]) {
				break
			} else {
				decoded = append([]byte{0x00}, decoded...)
			}
		}
	}

	return decoded
}
