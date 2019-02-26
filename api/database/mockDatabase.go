package database

import "github.com/kansuke231/go-with-vue/api/models"

type MockDB struct{}

func (mockDB *MockDB) UpdateNewsArticle(id int, rating int) {}

func (mockDB *MockDB) GetBestNews() []*models.NewsArticle {
	best := []*models.NewsArticle{}
	return best
}

func (mockDB *MockDB) GetAll() []*models.NewsArticle {
	all := []*models.NewsArticle{}
	first := &models.NewsArticle{
		ID:          1,
		Title:       "SomeTitle1",
		Link:        "SomeLink1",
		Description: "SomeDescription1",
		Rating:      3,
	}

	second := &models.NewsArticle{
		ID:          2,
		Title:       "SomeTitle2",
		Link:        "SomeLink2",
		Description: "SomeDescription2",
		Rating:      5,
	}
	all = append(all, first, second)
	return all
}
