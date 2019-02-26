package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kansuke231/go-with-vue/api/database"
	"github.com/kansuke231/go-with-vue/api/models"
)

func TestNewsFeedsHandler(t *testing.T) {

	context := &Context{DB: &database.MockDB{}, BestNews: nil}
	newsFeedsHandler := AppHandler{context, NewsFeedsHandler}
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/newsfeed", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := newsFeedsHandler

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := []*models.NewsArticle{
		&models.NewsArticle{
			ID:          1,
			Title:       "SomeTitle1",
			Link:        "SomeLink1",
			Description: "SomeDescription1",
			Rating:      3,
		},
		&models.NewsArticle{
			ID:          2,
			Title:       "SomeTitle2",
			Link:        "SomeLink2",
			Description: "SomeDescription2",
			Rating:      5,
		},
	}

	result := []*models.NewsArticle{}
	json.Unmarshal(rr.Body.Bytes(), &result)

	if len(result) != len(expected) {
		t.Errorf("The length of result array unmatched with that of expected. Length: %d", len(result))
	} else {

		for i := 0; i < 2; i++ {
			e := expected[i]
			r := result[i]

			if e.ID != r.ID {
				t.Errorf("IDs differ. Expected: %d  Returned: %d \n", e.ID, r.ID)
			}

			if e.Title != r.Title {
				t.Errorf("Titles differ. Expected: %s  Returned: %s \n", e.Title, r.Title)
			}

			if e.Link != r.Link {
				t.Errorf("Links differ. Expected: %s  Returned: %s \n", e.Link, r.Link)
			}

			if e.Description != r.Description {
				t.Errorf("Descriptions differ. Expected: %s  Returned: %s \n", e.Description, r.Description)
			}

			if e.Rating != r.Rating {
				t.Errorf("Ratings differ. Expected: %d  Returned: %d \n", e.Rating, r.Rating)
			}
		}
	}

}
