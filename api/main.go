package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kansuke231/go-with-vue/api/database"
	"github.com/kansuke231/go-with-vue/api/models"
)

func main() {

	println("Making the connection to the PostgreSQL instance.....")
	db_connection_string := "host=db port=5432 dbname=postgres user=docker password=docker sslmode=disable"
	time.Sleep(2 * time.Second)

	db, err := database.NewDB(db_connection_string)
	if err != nil {
		println(err.Error())
	}

	println("Connected!")
	test := &models.Test{ID: 13, SomeColumn: "some_value_13"}
	//db.CreateTable(test)
	println("db.HasTable(test) --> ", db.HasTable(test))
	//db.InsertTest(test)

	all := db.GetAll()

	for _, e := range all {
		println("---------")
		println(e.ID, e.SomeColumn)

	}

	println("Before defining router")
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/hello").Name("Hello").Handler(http.HandlerFunc(HelloHandler))
	println("After defining router")

	http.ListenAndServe(":8080", router)

	println("Done.")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("{this:'ok'}")
}
