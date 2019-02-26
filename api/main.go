package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kansuke231/go-with-vue/api/database"
	"github.com/kansuke231/go-with-vue/api/models"
)

type Context struct {
	DB       database.Database
	BestNews *models.BestNews
}

func main() {

	db := db_setup()
	best_news := generateBestNews(db)

	context := &Context{DB: db, BestNews: best_news}

	router := mux.NewRouter().StrictSlash(true)

	newsFeedsHandler := AppHandler{context, NewsFeedsHandler}
	router.Methods("GET").Path("/newsfeed").Name("newsfeed").Handler(newsFeedsHandler)

	updateHandler := AppHandler{context, UpdateHandler}
	router.Methods("PUT").Path("/newsfeed/{id}").Name("newsfeedUpdate").Handler(updateHandler)

	bestNewsHandler := AppHandler{context, BestNewsHandler}
	router.Methods("GET").Path("/bestnews").Name("bestnews").Handler(bestNewsHandler)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	// Schedule the update every 5 minutes.
	schedule(db, best_news, updateBestNews, 5*time.Minute)

	http.ListenAndServe(":8080", router)

}

func db_setup() *database.DB {
	db_connection_string := "host=db port=5432 dbname=postgres user=docker password=docker sslmode=disable"

	// Wait for 4 seconds to make sure Postgres instance is running and newsfeeder has inserted entries..
	time.Sleep(4 * time.Second)

	log.Println("Making the connection to the PostgreSQL instance.....")
	db, err := database.NewDB(db_connection_string)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
