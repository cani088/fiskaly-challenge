package domain

import (
	"encoding/base64"
	"errors"
	"fmt"
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
	LastSignature    string
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

func (d *Device) GetDecodedKeyPair() (interface{}, error) {
	if d.Algorithm == "RSA" {
		return d.GetDecodedRSAKeyPair()
	}

	if d.Algorithm == "ECC" {
		return d.GetDecodedECCKeyPair()
	}

	return nil, errors.New(fmt.Sprintf("Algorithm %s is not supported", d.Algorithm))
}

func (d *Device) GetDecodedECCKeyPair() (*crypto.ECCKeyPair, error) {
	marshaler := crypto.ECCMarshaler{}
	return marshaler.Decode(d.PrivateKey)
}

func (d *Device) GetDecodedRSAKeyPair() (*crypto.RSAKeyPair, error) {
	marshaler := crypto.RSAMarshaler{}
	return marshaler.Unmarshal(d.PrivateKey)
}

func (d *Device) SignData(data string) (signature string, signedData string) {
	signedData = string(d.SignatureCounter) + "_" + data + "_"
	if d.SignatureCounter > 0 {
		signedData += d.LastSignature
	} else {
		inputBytes := []byte(d.ID)
		signedData += base64.StdEncoding.EncodeToString(inputBytes)
	}

	if d.Algorithm == "RSA" {
		keyPair, _ := d.GetDecodedRSAKeyPair()
		signer := crypto.NewRSASigner(keyPair.Private)
		signature, _ := signer.Sign([]byte(signedData))
		return string(signature), signedData
	}

	if d.Algorithm == "RSA" {
		keyPair, _ := d.GetDecodedECCKeyPair()
		signer := crypto.NewECDSASigner(keyPair.Private)
		signature, err := signer.Sign([]byte(signedData))
		if err != nil {
			return "", ""
		}
		return string(signature), signedData
	}

	return "", signedData
}
