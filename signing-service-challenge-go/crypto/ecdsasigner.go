package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// ECDSASigner implements the Signer interface for ECDSA.
type ECDSASigner struct {
	privateKey *ecdsa.PrivateKey
}

// NewECDSASigner creates a new ECDSASigner with the provided private key.
func NewECDSASigner(privateKey *ecdsa.PrivateKey) *ECDSASigner {
	return &ECDSASigner{privateKey: privateKey}
}

// Sign implements the Sign method of the Signer interface for ECDSA.
func (s *ECDSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	hashed := sha256.Sum256(dataToBeSigned)
	r, a, err := ecdsa.Sign(rand.Reader, s.privateKey, hashed[:])
	if err != nil {
		return nil, err
	}

	// ECDSA signature is represented as a pair of integers (r, s)
	signature := append(r.Bytes(), a.Bytes()...)

	return signature, nil
}

// VerifyECDSASignature verifies an ECDSA signature.
func VerifyECDSASignature(publicKey *ecdsa.PublicKey, data []byte, signature []byte) bool {
	hashed := sha256.Sum256(data)

	// Extract r and s from the signature
	rBytes := signature[:len(signature)/2]
	sBytes := signature[len(signature)/2:]

	// Convert r and s back to big.Int
	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	return ecdsa.Verify(publicKey, hashed[:], r, s)
}
