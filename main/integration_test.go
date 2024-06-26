package main

import (
	
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPointsHandler(t *testing.T) {

	router := RegisterRoutes()

	receipt := Receipt{
		Retailer:     "Target",
		Total:        "35.35",
		Items:        []Item{{ShortDescription: "Mountain Dew 12PK", Price: "6.49"}, {ShortDescription: "Emils Cheese Pizza", Price: "12.25"}, {ShortDescription: "Knorr Creamy Chicken", Price: "1.26"}, {ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"}  },
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
	}

	body, err := json.Marshal(receipt)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var processResponse ProcessResponse
	err = json.NewDecoder(rr.Body).Decode(&processResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, processResponse.ID)

	req, err = http.NewRequest("GET", "/receipts/"+processResponse.ID+"/points", nil)
	assert.NoError(t, err)

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var pointsResponse PointsResponse
	err = json.NewDecoder(rr.Body).Decode(&pointsResponse)
	assert.NoError(t, err)
	assert.Equal(t, 28, pointsResponse.Points, "Expected points to be 28")

}
