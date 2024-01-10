package test

import (
	"bytes"
	"testing"

	algo "github.com/denizumutdereli/go-cryptography/algorithms"
)

// TestBase58EncodeDecode tests the Base58Encode and Base58Decode functions for correctness.
func TestBase58EncodeDecode(t *testing.T) {
	input := []byte("hello world")
	encoded := algo.Base58Encode(input)
	decoded := algo.Base58Decode(encoded)

	if !bytes.Equal(decoded, input) {
		t.Errorf("Base58Decode(Base58Encode(%s)) = %s; want %s", input, decoded, input)
	}
}
