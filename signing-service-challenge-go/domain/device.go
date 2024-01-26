package domain

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/api"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/google/uuid"
)

type Device struct {
	ID               string
	Label            string
	Algorithm        string
	SignatureCounter int
	RSAKeyPair       *crypto.RSAKeyPair
	ECCKeyPair       *crypto.ECCKeyPair
}

type DeviceRepository interface {
	AddDevice(device Device) (Device, error)
	GetDeviceById(id string) (Device, error)
	GetDeviceByLabel(label string) (Device, error)
	IncreaseDeviceCounter(id string) error
}

var server = api.Server{}
var storageService = server.StorageService

// Create is a method of the Domain struct that creates a new instance of Domain.
func (d *Device) Create(label string, algorithm string) (Device, error) {
	d.ID = uuid.NewString()
	d.Label = label
	d.Algorithm = algorithm
	d.SignatureCounter = 0

	if algorithm == "RSA" {
		generator := crypto.RSAGenerator{}
		d.RSAKeyPair, _ = generator.Generate()
	}

	if algorithm == "ECC" {
		generator := crypto.ECCGenerator{}
		d.ECCKeyPair, _ = generator.Generate()
	}

	device, err := storageService.AddDevice(*d)

	if err != nil {
		return Device{}, err
	}

	return device, err
}

func (d *Device) GetById(id string) (Device, error) {
	device, err := storageService.GetDeviceById(id)

	if err != nil {
		return Device{}, err
	}

	return device, nil
}
