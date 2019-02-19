package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kansuke231/go-with-vue/api/database"
)

func main() {

	println("Making the connection to the PostgreSQL instance.....")
	db_connection_string := "host=db port=5432 dbname=postgres user=docker password=docker sslmode=disable"
	time.Sleep(2 * time.Second)

	db, err := database.NewDB(db_connection_string)
	if err != nil {
		println(err.Error())
	}

	//test := &models.Test{ID: 13, SomeColumn: "some_value_13"}
	//db.CreateTable(test)
	//db.InsertTest(test)

	all := db.GetAll()

	for _, e := range all {
		println(e.ID, e.SomeColumn)

	}

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/hello").Name("Hello").Handler(http.HandlerFunc(HelloHandler))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Got a request!!!")
		http.ServeFile(w, r, "./index.html")
	})

	http.ListenAndServe(":8080", router)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	SetCrossOrigin(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("{\"id\": 13, \"some_column\": \"value\"}")
}

func SetCrossOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
