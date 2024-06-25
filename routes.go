package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() {

	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", ProcessReceiptsHandler).Methods("POST")

	http.Handle("/", r)
}
