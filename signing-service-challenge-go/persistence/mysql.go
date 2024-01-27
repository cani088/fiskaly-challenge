package persistence

import (
	"database/sql"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) AddDevice(device domain.Device) {
	//TODO: insert a device into MySQL
}

func (r *MySQLRepository) GetDeviceByID(id int) (domain.Device, error) {
	//TODO: retrieve a device by ID from MySQL
	return domain.Device{}, nil
}
