package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kansuke231/go-with-vue/api/database"
)

var static_result = generateStaticResults()

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/hello").Name("Hello").Handler(http.HandlerFunc(HelloHandler))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Got a request!!!")
		http.ServeFile(w, r, "./index.html")
	})

	schedule(static_result, updateStaticResult, 15*time.Second)

	http.ListenAndServe(":8080", router)

}

type StaticResult struct {
	ID         int    `json:"id"`
	SomeResult string `json:"some_result"`
}

func schedule(s *StaticResult, update func(s *StaticResult), delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			// Every delay (e.g. 5 minutes), update() gets executed.
			println("In go func()")
			select {
			case <-time.After(delay):
				// Do update here.
				println("Updating......")
				update(s)

			}
		}
	}()

	return stop
}

func updateStaticResult(s *StaticResult) {
	s.ID = rand.Int()
	s.SomeResult = "ThisIsSomeResult" + strconv.Itoa(rand.Int())
}

func generateStaticResults() *StaticResult {
	return &StaticResult{ID: rand.Int(), SomeResult: "ThisIsSomeResult" + strconv.Itoa(rand.Int())}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes

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

	SetCrossOrigin(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	println(static_result.ID)

	json.NewEncoder(w).Encode(static_result)
}

func SetCrossOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
