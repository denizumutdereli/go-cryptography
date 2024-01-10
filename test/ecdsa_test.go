package test

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	algo "github.com/denizumutdereli/go-cryptography/algorithms"
)

// TestGenerateECDSAKeys tests the ECDSA key generation and signature functionality.
func TestGenerateECDSAKeys(t *testing.T) {
	privateKey, err := algo.GenerateECDSAKeys()
	if err != nil {
		t.Fatalf("Failed to generate ECDSA keys: %v", err)
	}

	// Test signature generation and verification
	message := "test message"
	hash := sha256.Sum256([]byte(message))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		t.Fatalf("Failed to sign message: %v", err)
	}

	valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	if !valid {
		t.Errorf("ECDSA signature verification failed")
	}

	// Test private key to public address conversion
	publicKeyBytes := crypto.FromECDSAPub(&privateKey.PublicKey)
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	if len(publicKeyBytes) == 0 || address.String() == "" {
		t.Errorf("Failed to derive public key or address from private key")
	}
}
