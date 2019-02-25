package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func NewsFeedsHandler(db *database.DB, w http.ResponseWriter, r *http.Request) {

	all := db.GetAll()

	SetCrossOrigin(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(all)
}

func UpdateHandler(db *database.DB, w http.ResponseWriter, r *http.Request) {
	println("In UpdateHandler ")
	SetCrossOrigin(w)
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)

	if err != nil || idInt < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("{\"error:\" \"invalid parameter\"}")
		return
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	rating, _ := strconv.Atoi(string(body))
	println("Before UpdateNewsArticle ")
	db.UpdateNewsArticle(idInt, rating)
	println("After UpdateNewsArticle ")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	//json.NewEncoder(w).Encode("")
}

func SetCrossOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
