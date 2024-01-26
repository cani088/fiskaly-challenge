package persistence

import (
	"errors"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/api"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type InMemoryRepository struct {
	devices map[int]domain.Device
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (m *InMemoryRepository) AddDevice(device domain.Device) (domain.Device, error) {
	_, err := m.GetDeviceByLabel(device.Label)

	if err == nil {
		_ = append(api.Devices, device)
		return device, nil
	}

	return domain.Device{}, err
}

func (m *InMemoryRepository) GetDeviceById(id string) (domain.Device, error) {
	for i := range api.Devices {
		if api.Devices[i].ID == id {
			return api.Devices[i], nil
		}
	}

	return domain.Device{}, errors.New("device with id: " + id + " does not exist")
}

func (m *InMemoryRepository) GetDeviceByLabel(label string) (interface{}, error) {
	for i := range api.Devices {
		if api.Devices[i].Label == label {
			return api.Devices[i], nil
		}
	}

	return nil, errors.New("device with label: " + label + " does not exist")
}

func (m *InMemoryRepository) IncreaseDeviceCounter(id string) error {
	for i := range api.Devices {
		if api.Devices[i].ID == id {
			api.Devices[i].SignatureCounter = api.Devices[i].SignatureCounter + 1
			return nil
		}
	}

	return errors.New("Device could not be found")
}
