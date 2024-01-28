package domain

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
)

type Signature struct {
	Signature  string
	SignedData string
	Device     Device
}

func NewSignature(signature string, signedData string, device Device) *Signature {
	return &Signature{
		Signature:  signature,
		SignedData: signedData,
		Device:     device,
	}
}

func (s *Signature) Verify() bool {
	if s.Device.Algorithm == "ECC" {
		var keyPair, _ = s.Device.GetDecodedECCKeyPair()
		verified := crypto.VerifyECDSASignature(keyPair.Public, []byte(s.Signature), []byte(s.SignedData))
		if verified {
			return true
		}
	}

	if s.Device.Algorithm == "RSA" {
		var keyPair, _ = s.Device.GetDecodedRSAKeyPair()
		_, err := crypto.VerifyRSASignature(keyPair.Public, []byte(s.Signature), []byte(s.SignedData))
		if err == nil {
			return true
		}
	}

	return false
}
