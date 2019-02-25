package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Custom handler type so that each handler for an endpoint has DB dependency injection.
type AppHandler struct {
	context *Context
	H       func(*Context, http.ResponseWriter, *http.Request)
}

func (handler AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Updated to pass appContext as a parameter to our handler type.
	handler.H(handler.context, w, r)
}

func NewsFeedsHandler(context *Context, w http.ResponseWriter, r *http.Request) {

	all := context.DB.GetAll()

	SetCrossOrigin(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(all)
}

func BestNewsHandler(context *Context, w http.ResponseWriter, r *http.Request) {

	SetCrossOrigin(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(context.BestNews)
}

func UpdateHandler(context *Context, w http.ResponseWriter, r *http.Request) {
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
	context.DB.UpdateNewsArticle(idInt, rating)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	//json.NewEncoder(w).Encode("")
}

func SetCrossOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
