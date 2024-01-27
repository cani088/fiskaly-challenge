package domain

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
)

type Device struct {
	ID               string
	Label            string
	Algorithm        string
	SignatureCounter int
	RSAKeyPair       *crypto.RSAKeyPair
	ECCKeyPair       *crypto.ECCKeyPair
}
