package main

import (
	"time"

	"github.com/kansuke231/go-with-vue/api/database"
	"github.com/kansuke231/go-with-vue/api/models"
)

// function schedule spawns a go routine that updates BestNews every delay specified.
func schedule(db *database.DB, bestNews *models.BestNews, update func(db *database.DB, s *models.BestNews), delay time.Duration) {
	go func() {
		for {
			// Every delay (e.g. 5 minutes), update() gets executed.
			select {
			case <-time.After(delay):
				update(db, bestNews)

			}
		}
	}()
}

func updateBestNews(db *database.DB, bestNews *models.BestNews) {
	bestNews.TopRated = db.GetBestNews()
	bestNews.Created = time.Now().String()
}

func generateBestNews(db *database.DB) *models.BestNews {
	return &models.BestNews{TopRated: db.GetBestNews(), Created: time.Now().String()}
}
