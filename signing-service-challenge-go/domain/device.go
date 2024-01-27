package domain

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/google/uuid"
)

type Device struct {
	ID               string
	Label            string
	Algorithm        string
	SignatureCounter int
	PrivateKey       []byte
	PublicKey        []byte
}

func NewDevice(label string, algorithm string) *Device {
	device := &Device{
		ID:               uuid.NewString(),
		Label:            label,
		Algorithm:        algorithm,
		SignatureCounter: 0,
	}
	device.GenerateKeys()
	return device
}

func (d *Device) GenerateKeys() {
	// TODO: try to make this more elegant
	if d.Algorithm == "RSA" {
		generator := crypto.RSAGenerator{}
		keyPair, _ := generator.Generate()
		rsaMarshaler := crypto.NewRSAMarshaler()
		d.PrivateKey, d.PublicKey, _ = rsaMarshaler.Marshal(*keyPair)
	}

	if d.Algorithm == "ECC" {
		generator := crypto.ECCGenerator{}
		keyPair, _ := generator.Generate()
		eccMarshaler := crypto.NewECCMarshaler()
		d.PrivateKey, d.PublicKey, _ = eccMarshaler.Encode(*keyPair)
	}
}
