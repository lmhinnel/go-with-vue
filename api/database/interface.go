package database

import "github.com/kansuke231/go-with-vue/api/models"

type Database interface {
	UpdateNewsArticle(id int, rating int)
	GetBestNews() []*models.NewsArticle
	GetAll() []*models.NewsArticle
}
