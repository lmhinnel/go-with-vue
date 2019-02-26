package database

import "github.com/kansuke231/go-with-vue/api/models"

type Database interface {
	HasTable(test *models.NewsArticle) bool
	InsertTest(test *models.NewsArticle)
	UpdateNewsArticle(id int, rating int)
	GetBestNews() []*models.NewsArticle
	GetAll() []*models.NewsArticle
}
