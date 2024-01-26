package persistence

import (
	"errors"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type Mysql struct{}

var naMessage = "mysql service not available"

func (db *Mysql) AddDevice(device domain.Device) (domain.Device, error) {
	return domain.Device{}, errors.New(naMessage)
}

func (db *Mysql) GetDeviceById(id string) (domain.Device, error) {
	return domain.Device{}, errors.New(naMessage)
}

func (db *Mysql) GetDeviceByLabel(label string) (interface{}, error) {
	return nil, errors.New(naMessage)
}

func (db *Mysql) IncreaseDeviceCounter(id string) error {
	return errors.New(naMessage)
}
