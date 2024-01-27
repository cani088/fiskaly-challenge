package persistence

import "github.com/fiskaly/coding-challenges/signing-service-challenge/domain"

type StorageInterface interface {
	AddDevice(device domain.Device) error
	GetDeviceByLabel(label string) (domain.Device, error)
	IncreaseDeviceCounter(label string) error
	UpdateLastSignature(label string, signature string) error
}
