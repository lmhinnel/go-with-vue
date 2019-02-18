package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Struct DB implements Database interface.
type DB struct {
	database *gorm.DB
}

func NewDB(connectionString string) (*DB, error) {
	transdb, err := gorm.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return &DB{transdb}, nil
}
func (db *DB) IsAlive() bool {
	return true
}
