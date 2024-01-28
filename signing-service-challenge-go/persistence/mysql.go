package persistence

import (
	"database/sql"
	"errors"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

var e = errors.New("mysql storage not available")

func NewMySQLRepository(db *sql.DB) (*MySQLRepository, error) {
	return &MySQLRepository{
		db: db,
	}, e
}

func (r *MySQLRepository) AddDevice(device domain.Device) error {
	return e
}

func (r *MySQLRepository) GetDeviceByLabel(label string) (domain.Device, error) {
	return domain.Device{}, e
}

func (r *MySQLRepository) IncreaseDeviceCounter(label string) error {
	return e
}

func (r *MySQLRepository) UpdateLastSignature(label string, signature string) error {
	return e
}

func (r *MySQLRepository) GetAllDevices() error {
	return e
}
