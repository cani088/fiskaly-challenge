package persistence

import (
	"errors"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type InMemoryRepository struct {
	devices map[string]domain.Device
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		devices: make(map[string]domain.Device),
	}
}

func (m *InMemoryRepository) AddDevice(device domain.Device) {
	m.devices[device.ID] = device
}

func (m *InMemoryRepository) GetDeviceById(id string) (domain.Device, interface{}) {
	device, ok := m.devices[id]
	return device, ok
}

func (m *InMemoryRepository) IncreaseDeviceCounter(id string) (domain.Device, error) {
	var device = m.devices[id]
	if device.ID == "" {
		return domain.Device{}, errors.New("Device does not exist")
	}
	device.SignatureCounter = device.SignatureCounter + 1
	m.devices[id] = device
	return m.devices[id], nil
}
