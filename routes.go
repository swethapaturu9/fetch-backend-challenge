package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", ProcessReceiptsHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", GetPointsHandler).Methods("GET")

	http.Handle("/", r)

	return r 
}
