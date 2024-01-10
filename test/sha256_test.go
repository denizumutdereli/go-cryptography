package test

import (
	"testing"

	algo "github.com/denizumutdereli/go-cryptography/algorithms"
)

// TestSha256Hash tests the Sha256Hash function with a basic string.
// It verifies that the SHA-256 hash is computed correctly.
// https://codebeautify.org/sha256-hash-generator
func TestSha256Hash(t *testing.T) {
	input := "hello world"
	expectedOutput := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	actualOutput := algo.Sha256Hash(input)
	if actualOutput != expectedOutput {
		t.Errorf("Sha256Hash(%s) = %s; want %s", input, actualOutput, expectedOutput)
	}
}

// TestSha3Hash tests the Sha3Hash function with a basic string.
// It verifies that the SHA3-256 hash is computed correctly.
// https://codebeautify.org/sha3-256-hash-generator
func TestSha3Hash(t *testing.T) {
	input := "hello world"
	expectedOutput := "644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938"
	actualOutput := algo.Sha3Hash(input)
	if actualOutput != expectedOutput {
		t.Errorf("Sha3Hash(%s) = %s; want %s", input, actualOutput, expectedOutput)
	}
}
