package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kansuke231/go-with-vue/api/database"
)

var static_result = generateStaticResults()

func main() {

	db := db_setup()

	router := mux.NewRouter().StrictSlash(true)

	newsFeedsHandler := AppHandler{db, NewsFeedsHandler}
	router.Methods("GET").Path("/newsfeed").Name("newsfeed").Handler(newsFeedsHandler)

	updateHandler := AppHandler{db, UpdateHandler}
	router.Methods("PUT").Path("/newsfeed/{id}").Name("newsfeedUpdate").Handler(updateHandler)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	// Schedule the update every 5 seconds.
	schedule(static_result, updateStaticResult, 5*time.Second)

	http.ListenAndServe(":8080", router)

}

func db_setup() *database.DB {
	db_connection_string := "host=db port=5432 dbname=postgres user=docker password=docker sslmode=disable"

	// Wait for 2 seconds to make sure Postgres instance is running.
	time.Sleep(2 * time.Second)

	log.Println("Making the connection to the PostgreSQL instance.....")
	db, err := database.NewDB(db_connection_string)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
