package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mmcdole/gofeed"
)

type NewsArticle struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Rating      int    `json:"rating"` // This property does not exist in the original news feed.
}

func NewDB(connectionString string) (*gorm.DB, error) {
	transdb, err := gorm.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return transdb, nil
}

func main() {

	fp := gofeed.NewParser()
	techFeed, _ := fp.ParseURL("http://rss.nytimes.com/services/xml/rss/nyt/Technology.xml")
	europeFeed, _ := fp.ParseURL("http://rss.nytimes.com/services/xml/rss/nyt/Europe.xml")

	feeds := []*NewsArticle{}

	for _, item := range techFeed.Items {
		article := &NewsArticle{Title: item.Title, Link: item.Link, Description: item.Description}
		feeds = append(feeds, article)
	}

	for _, item := range europeFeed.Items {
		article := &NewsArticle{Title: item.Title, Link: item.Link, Description: item.Description}
		feeds = append(feeds, article)
	}

	// Wait for 2 seconds to make sure Postgres instance is running.
	time.Sleep(2 * time.Second)

	// Inserting the feeds into PostgreSQL
	db_connection_string := "host=db port=5432 dbname=postgres user=docker password=docker sslmode=disable"
	db, _ := NewDB(db_connection_string)

	// Create table.
	db.AutoMigrate(&NewsArticle{})

	for _, feed := range feeds {
		db.Create(feed)
	}

	// Let's see if they were inserted...
	all := []*NewsArticle{}
	db.Where("true").Find(&all)

	for _, feed := range all {
		println(feed.ID, feed.Title, feed.Rating)
	}

}
