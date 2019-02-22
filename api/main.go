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

	// test1 := &models.NewsArticle{ID: 13, SomeColumn: "some_value_13", Rating: 5}
	// println("Creating news_articles table...")
	// db.CreateTable(test1)
	// db.InsertTest(test1)
	//
	// test2 := &models.NewsArticle{ID: 15, SomeColumn: "some_value_15", Rating: 2}
	// db.InsertTest(test2)

	router := mux.NewRouter().StrictSlash(true)

	helloHandler := AppHandler{db, HelloHandler}
	router.Methods("GET").Path("/hello").Name("Hello").Handler(helloHandler)

	updateHandler := AppHandler{db, UpdateHandler}
	router.Methods("PUT").Path("/{id}").Name("Update").Handler(updateHandler)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	// Schedule the update every 5 seconds.
	schedule(static_result, updateStaticResult, 5*time.Second)

	http.ListenAndServe(":8080", router)

}

func db_setup() *database.DB {

	log.Println("Making the connection to the PostgreSQL instance.....")
	db_connection_string := "host=db port=5432 dbname=postgres user=docker password=docker sslmode=disable"

	// Wait for 2 seconds to make sure Postgres instance is running.
	time.Sleep(2 * time.Second)

	db, err := database.NewDB(db_connection_string)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
