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

var availableEncryptionTypes = []string{"RSA", "ECC"}

func NewDevice(label string, algorithm string) (*Device, error) {
	if label == "" {
		return nil, errors.New("label cannot be empty")
	}

	if !isInArray(algorithm, availableEncryptionTypes) {
		return nil, errors.New("encryption method is not available")
	}

	device := &Device{
		ID:               uuid.NewString(),
		Label:            label,
		Algorithm:        algorithm,
		SignatureCounter: 0,
	}
	device.GenerateKeys()

	return device, nil
}

func isInArray(needle any, haystack []string) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}
	return false
}

func (d *Device) GenerateKeys() {
	if d.Algorithm == "RSA" {
		generator := crypto.RSAGenerator{}
		keyPair, _ := generator.Generate()
		rsaMarshaler := crypto.NewRSAMarshaler()
		d.PublicKey, d.PrivateKey, _ = rsaMarshaler.Marshal(*keyPair)
	}

	if d.Algorithm == "ECC" {
		generator := crypto.ECCGenerator{}
		keyPair, err := generator.Generate()
		if err != nil {
			print(err)
		}
		eccMarshaler := crypto.NewECCMarshaler()
		d.PublicKey, d.PrivateKey, _ = eccMarshaler.Encode(*keyPair)
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
	keyPair, err := marshaler.Unmarshal(d.PrivateKey)
	return keyPair, err
}

func (d *Device) SignData(data string) (signature string, signedData string) {
	signedData = fmt.Sprintf("%d_%s_", d.SignatureCounter, data)
	inputBytes := []byte(d.ID)

	if d.SignatureCounter > 0 {
		inputBytes = []byte(d.LastSignature)
	}
	signedData += base64.StdEncoding.EncodeToString(inputBytes)

	if d.Algorithm == "RSA" {
		keyPair, _ := d.GetDecodedRSAKeyPair()
		signer := crypto.NewRSASigner(keyPair.Private)
		signature, err := signer.Sign([]byte(signedData))
		if err != nil {
			return "", ""
		}
		return base64.StdEncoding.EncodeToString(signature), signedData
	}

	if d.Algorithm == "ECC" {
		keyPair, _ := d.GetDecodedECCKeyPair()
		signer := crypto.NewECDSASigner(keyPair.Private)
		signature, err := signer.Sign([]byte(signedData))
		if err != nil {
			return "", ""
		}
		return base64.StdEncoding.EncodeToString(signature), signedData
	}

	return "", signedData
}
