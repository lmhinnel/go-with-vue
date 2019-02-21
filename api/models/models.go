package models

type NewsArticle struct {
	ID         int    `json:"id"`
	SomeColumn string `json:"some_column"`
	Rating     int    `json:"rating"`
}

type StaticResult struct {
	ID         int    `json:"id"`
	SomeResult string `json:"some_result"`
}
