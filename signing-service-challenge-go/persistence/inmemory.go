package persistence

import (
	"errors"
	"fmt"
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

func (m *InMemoryRepository) AddDevice(device domain.Device) error {
	for _, value := range m.devices {
		if device.Label == value.Label {
			return errors.New(fmt.Sprintf("Device with label '%s' already exists", device.Label))
		}
	}
	m.devices[device.ID] = device
	return nil
}

func (m *InMemoryRepository) GetDeviceById(id string) (domain.Device, error) {
	device, _ := m.devices[id]
	if device.ID == "" {
		return domain.Device{}, errors.New(fmt.Sprintf("Device with label '%s' already exists", device.Label))
	}
	return device, nil
}

func (m *InMemoryRepository) IncreaseDeviceCounter(id string) (domain.Device, error) {
	var device = m.devices[id]
	if device.ID == "" {
		return domain.Device{}, errors.New("device does not exist")
	}
	device.SignatureCounter = device.SignatureCounter + 1
	m.devices[id] = device
	return m.devices[id], nil
}
