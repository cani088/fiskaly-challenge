package persistence

import (
	"errors"
	"fmt"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"strconv"
	"sync"
)

type InMemoryRepository struct {
	devices    map[string]domain.Device
	signatures map[string]domain.Signature
}

var mutex sync.Mutex

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		devices:    make(map[string]domain.Device),
		signatures: make(map[string]domain.Signature),
	}
}

func (m *InMemoryRepository) AddDevice(device domain.Device) error {
	mutex.Lock()
	defer mutex.Unlock()

	// Check if device with the Label already exists
	for i := range m.devices {
		if m.devices[i].Label == device.Label {
			return errors.New(fmt.Sprintf("Device with label '%s' already exists", device.Label))
		}
	}

	m.devices[device.ID] = device
	return nil
}

func (m *InMemoryRepository) GetDeviceById(id string) (domain.Device, error) {
	mutex.Lock()
	defer mutex.Unlock()
	device, _ := m.devices[id]
	if device.ID == "" {
		return domain.Device{}, errors.New(fmt.Sprintf("Device with id '%s' does not exist", device.ID))
	}
	return device, nil
}

func (m *InMemoryRepository) IncreaseDeviceCounter(id string) error {
	mutex.Lock()
	defer mutex.Unlock()
	var device = m.devices[id]
	if device.ID == "" {
		return errors.New("device does not exist")
	}
	device.SignatureCounter = device.SignatureCounter + 1
	m.devices[id] = device
	return nil
}

func (m *InMemoryRepository) UpdateLastSignature(id string, signature string) error {
	mutex.Lock()
	defer mutex.Unlock()
	var device = m.devices[id]
	if device.ID == "" {
		return errors.New("device does not exist")
	}
	device.LastSignature = signature
	m.devices[id] = device
	return nil
}

func (m *InMemoryRepository) GetAllDevices() any {
	mutex.Lock()
	defer mutex.Unlock()
	var modifiedList []map[string]string

	for _, item := range m.devices {
		modifiedItem := map[string]string{
			"id":               item.ID,
			"label":            item.Label,
			"signatureCounter": strconv.Itoa(item.SignatureCounter),
		}

		modifiedList = append(modifiedList, modifiedItem)
	}
	return modifiedList
}
