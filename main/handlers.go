package main

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var (
	pointsMap                     = make(map[string]int)
	validate  *validator.Validate = validator.New()
)

func init() {

	RegisterValidations(validate)
}

type ProcessResponse struct {
	ID string `json:"id"`
}

func ProcessReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var receipt Receipt

	if err := json.Unmarshal(body, &receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(receipt); err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	id := generateUniqueID()

	points := CalculatePoints(receipt)

	pointsMap[id] = points

	response := ProcessResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "The receipt is invalid", http.StatusInternalServerError)
	}

}

type PointsResponse struct {
	Points int `json:"points"`
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {

	if r.ContentLength > 0 {
        http.Error(w, "Payload not allowed", http.StatusBadRequest)
        return
    }

	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := pointsMap[id]
	if !exists {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	response := PointsResponse{
		Points: points,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
