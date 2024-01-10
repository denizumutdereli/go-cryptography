package test

import (
	"testing"

	algo "github.com/denizumutdereli/go-cryptography/algorithms"
)

// TestGenerateRIPEMD160Hash tests the GenerateRIPEMD160Hash function to ensure it generates the hash correctly.
// https://hash.rfctools.com/ripemd160-hash-generator/
func TestGenerateRIPEMD160Hash(t *testing.T) {
	input := "hello world"
	expectedOutput := "98c615784ccb5fe5936fbc0cbe9dfdb408d92f0f"
	actualOutput := algo.GenerateRIPEMD160Hash(input)
	if actualOutput != expectedOutput {
		t.Errorf("GenerateRIPEMD160Hash(%s) = %s; want %s", input, actualOutput, expectedOutput)
	}
}
