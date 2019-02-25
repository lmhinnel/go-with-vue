package models

type NewsArticle struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Rating      int    `json:"rating"` // This property does not exist in the original news feed.
}

type BestNews struct {
	TopRated []*NewsArticle `json:"top_rated_news"`
	Created  string         `json:"created_at"`
}
