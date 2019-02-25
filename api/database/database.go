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

func (db *DB) CreateTable(test *models.NewsArticle) {
	db.database.AutoMigrate(test)
}

func (db *DB) HasTable(test *models.NewsArticle) bool {
	return db.database.HasTable(test)
}

func (db *DB) InsertTest(test *models.NewsArticle) {
	db.database.Create(test)
}

func (db *DB) UpdateNewsArticle(id int, rating int) {
	println("In UpdateNewsArticle ", id, rating)
	article := &models.NewsArticle{ID: id}
	db.database.Model(article).Update("rating", rating)
}

func (db *DB) GetBestNews() []*models.NewsArticle {
	best := []*models.NewsArticle{}
	db.database.Order("rating desc").Limit(5).Find(&best)
	return best
}

func (db *DB) GetAll() []*models.NewsArticle {
	all := []*models.NewsArticle{}
	db.database.Where("true").Find(&all)
	return all
}
