package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kansuke231/go-with-vue/api/models"
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

func (db *DB) CreateTable(test *models.Test) {
	db.database.AutoMigrate(test)
}

func (db *DB) HasTable(test *models.Test) bool {
	return db.database.HasTable(test)
}

func (db *DB) InsertTest(test *models.Test) {
	db.database.Create(test)
}

func (db *DB) GetAll() []*models.Test {
	all := []*models.Test{}
	db.database.Where("true").Find(&all)
	return all
}
