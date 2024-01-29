package persistence

import "github.com/fiskaly/coding-challenges/signing-service-challenge/domain"

type StorageInterface interface {
	AddDevice(device domain.Device) error
	GetDeviceById(id string) (domain.Device, error)
	IncreaseDeviceCounter(id string) error
	UpdateLastSignature(id string, signature string) error
	GetAllDevices() any
	GetAllTransactions() any
	AddTransaction(transaction domain.Transaction) error
}
