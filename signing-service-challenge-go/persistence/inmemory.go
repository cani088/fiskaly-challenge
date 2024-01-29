package persistence

import (
	"errors"
	"fmt"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"strconv"
	"sync"
)

type InMemoryRepository struct {
	devices          map[string]domain.Device
	deviceLock       sync.RWMutex
	transactions     map[string]domain.Transaction
	transactionsLock sync.RWMutex
}

var mutex sync.Mutex

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		devices:          make(map[string]domain.Device),
		deviceLock:       sync.RWMutex{},
		transactions:     make(map[string]domain.Transaction),
		transactionsLock: sync.RWMutex{},
	}
}

func (m *InMemoryRepository) AddDevice(device domain.Device) error {
	m.deviceLock.Lock()
	defer m.deviceLock.Unlock()

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
	m.deviceLock.Lock()
	defer m.deviceLock.Unlock()
	device, _ := m.devices[id]
	if device.ID == "" {
		return domain.Device{}, errors.New(fmt.Sprintf("Device with id '%s' does not exist", device.ID))
	}
	return device, nil
}

func (m *InMemoryRepository) IncreaseDeviceCounter(id string) error {
	m.deviceLock.Lock()
	defer m.deviceLock.Unlock()
	var device = m.devices[id]
	if device.ID == "" {
		return errors.New("device does not exist")
	}
	device.SignatureCounter = device.SignatureCounter + 1
	m.devices[id] = device
	return nil
}

func (m *InMemoryRepository) UpdateLastSignature(id string, signature string) error {
	m.deviceLock.Lock()
	defer m.deviceLock.Unlock()
	var device = m.devices[id]
	if device.ID == "" {
		return errors.New("device does not exist")
	}
	device.LastSignature = signature
	m.devices[id] = device
	return nil
}

func (m *InMemoryRepository) GetAllDevices() any {
	m.deviceLock.Lock()
	defer m.deviceLock.Unlock()
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

func (m *InMemoryRepository) AddTransaction(transaction domain.Transaction) error {
	m.transactionsLock.Lock()
	defer m.transactionsLock.Unlock()

	m.transactions[transaction.ID] = transaction

	return nil
}

func (m *InMemoryRepository) GetAllTransactions() any {
	m.transactionsLock.Lock()
	defer m.transactionsLock.Unlock()
	var modifiedList []map[string]string

	for _, item := range m.transactions {
		modifiedItem := map[string]string{
			"transactionId": item.ID,
			"signature":     item.Signature,
			"signedData":    item.SignedData,
			"deviceId":      item.Device.ID,
		}

		modifiedList = append(modifiedList, modifiedItem)
	}
	return modifiedList
}
