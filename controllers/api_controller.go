package controllers

import (
	"net/http"
	"encoding/json"
	"strings"
	"github.com/gorilla/mux"
	"goipmserver/services"
)

type TSearch struct {
	Skip int
	Limit int
	Query string
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func get(w http.ResponseWriter, collection string, query interface{}, skip int, limit int) {
	searchResults, searchErrors := services.SearchCollection(collection, query, skip, limit)
	if searchErrors != "" {
		respondWithError(w, http.StatusInternalServerError, searchErrors)
	}
	respondWithJSON(w,http.StatusOK, searchResults )
}

func GetQueryHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	collection := vars["collection"]
	squery := vars["squery"]

	// create a json from the query as string
	var search TSearch
	var err error
	err = json.Unmarshal([]byte(squery), &search)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if search.Limit == 0 {
		search.Limit = 100
	}

	// transform bson.M for mgo
	var query interface{}
	search.Query = strings.Replace(search.Query, "'", "\"",-1)
	err = json.Unmarshal([]byte(search.Query), &query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	get(w, collection, query, search.Skip, search.Limit)
}

func GetHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	collection := vars["collection"]
	get(w, collection,nil,0,100)
}

