package persistence

import (
	"errors"
	"fmt"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"sync"
)

type InMemoryRepository struct {
	devices map[string]domain.Device
}

var mutex sync.Mutex

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		devices: make(map[string]domain.Device),
	}
}

func (m *InMemoryRepository) AddDevice(device domain.Device) error {
	mutex.Lock()
	defer mutex.Unlock()

	memoryDevice, _ := m.devices[device.Label]
	if memoryDevice.ID == "" {
		m.devices[device.Label] = device
		return nil
	}

	return errors.New(fmt.Sprintf("Device with label '%s' already exists", device.Label))
}

func (m *InMemoryRepository) GetDeviceByLabel(label string) (domain.Device, error) {
	mutex.Lock()
	defer mutex.Unlock()
	device, _ := m.devices[label]
	if device.ID == "" {
		return domain.Device{}, errors.New(fmt.Sprintf("Device with label '%s' does not exist", device.Label))
	}
	return device, nil
}

func (m *InMemoryRepository) IncreaseDeviceCounter(label string) error {
	mutex.Lock()
	defer mutex.Unlock()
	var device = m.devices[label]
	if device.ID == "" {
		return errors.New("device does not exist")
	}
	device.SignatureCounter = device.SignatureCounter + 1
	m.devices[label] = device
	return nil
}

func (m *InMemoryRepository) UpdateLastSignature(label string, signature string) error {
	mutex.Lock()
	defer mutex.Unlock()
	var device = m.devices[label]
	if device.ID == "" {
		return errors.New("device does not exist")
	}
	device.LastSignature = signature
	m.devices[label] = device
	return nil
}
