// Package algorithms provides various cryptographic algorithm implementations.
// This file includes the implementation for the ECDSA key generation using the secp256k1 curve.

package algorithms

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// GenerateECDSAKeys generates a new public and private key pair.
func GenerateECDSAKeys() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
}

// SignMessage signs a message using the given private key.
func SignMessage(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, message)
	if err != nil {
		return nil, err
	}

	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}
