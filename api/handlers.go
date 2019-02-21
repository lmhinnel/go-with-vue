package main

import (
	"encoding/json"
	"net/http"

	"github.com/kansuke231/go-with-vue/api/database"
)

// Custom handler type so that each handler for an endpoint has DB dependency injection.
type AppHandler struct {
	*database.DB
	H func(*database.DB, http.ResponseWriter, *http.Request)
}

func (handler AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Updated to pass appContext as a parameter to our handler type.
	handler.H(handler.DB, w, r)
}

func HelloHandler(db *database.DB, w http.ResponseWriter, r *http.Request) {

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

	json.NewEncoder(w).Encode(static_result)
}

func SetCrossOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
