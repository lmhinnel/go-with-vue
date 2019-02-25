package main

import (
	"time"

	"github.com/kansuke231/go-with-vue/api/database"
	"github.com/kansuke231/go-with-vue/api/models"
)

// function schedule spawns a go routine that updates StaticResult every delay specified.
func schedule(db *database.DB, bestNews *models.BestNews, update func(db *database.DB, s *models.BestNews), delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			// Every delay (e.g. 5 minutes), update() gets executed.
			select {
			case <-time.After(delay):
				// Do update here.
				update(db, bestNews)

			}
		}
	}()

	return stop
}

func updateBestNews(db *database.DB, bestNews *models.BestNews) {
	bestNews.TopRated = db.GetBestNews()
	bestNews.Created = time.Now().String()
}

func generateBestNews(db *database.DB) *models.BestNews {
	return &models.BestNews{TopRated: db.GetBestNews(), Created: time.Now().String()}
}
