package main

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/go-playground/validator/v10"
	
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
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	id := generateUniqueID()

	response := ProcessResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
