package test

import (
	"testing"

	algo "github.com/denizumutdereli/go-cryptography/algorithms"
)

// TestKeccak256Hash tests the Keccak256Hash function with a basic string.
// It verifies that the Keccak-256 hash is computed correctly.
// https://keccak-256.4tools.net/
func TestKeccak256Hash(t *testing.T) {
	input := "hello world"
	expectedOutput := "47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad"
	actualOutput := algo.Keccak256Hash(input)
	if actualOutput != expectedOutput {
		t.Errorf("Keccak256Hash(%s) = %s; want %s", input, actualOutput, expectedOutput)
	}
}
