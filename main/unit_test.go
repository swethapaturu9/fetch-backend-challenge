package main

import (
	"testing"
)

func TestRetailerNamePointsCalculator(t *testing.T) {
	calculator := &RetailerNamePointsCalculator{}

	// Positive test case
	receipt := Receipt{Retailer: "Target"}
	points := calculator.CalculatePoints(receipt)
	expected := 6
	if points != expected {
		t.Errorf("expected %d points, got %d", expected, points)
	}

	// Special characters
	receipt = Receipt{Retailer: "M&M Corner Market"}
	points = calculator.CalculatePoints(receipt)
	expected = 14 
	if points != expected {
		t.Errorf("expected %d points, got %d", expected, points)
	}
}

func TestItemDescriptionPointsCalculator(t *testing.T) {
	calculator := &ItemDescriptionPointsCalculator{}

	// characters multiple of 3 
	receipt := Receipt{Items: []Item{{ShortDescription: "Emils Cheese Pizza", Price: "12.25"}, {ShortDescription: "  Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"}}}
	points := calculator.CalculatePoints(receipt)
	expected := 6
	if points != expected {
		t.Errorf("expected %d points, got %d", expected, points)
	}


	//characters not multiple of 3 
	receipt = Receipt{Items: []Item{{ShortDescription: "Mountain Dew 12PK", Price: "6.49"}, {ShortDescription: "Knorr Creamy Chicken", Price: "1.26"}}}
	points = calculator.CalculatePoints(receipt)
	expected = 0
	if points != expected {
		t.Errorf("expected %d points, got %d", expected, points)
	}

}

func TestItemsPointsCalculator(t *testing.T) {
	calculator := &ItemsPointsCalculator{}

	//4 items 
	receipt := Receipt{Items: []Item{{ShortDescription: "Mountain Dew 12PK", Price: "6.49"}, {ShortDescription: "Emils Cheese Pizza", Price: "12.25"}, {ShortDescription: "Doritos Nacho Cheese", Price: "3.35"}, {ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"} }}
	points := calculator.CalculatePoints(receipt)
	expected := 10
	if points != expected {
		t.Errorf("expected %d points, got %d", expected, points)
	}

	//1 item
	receipt = Receipt{ Items: []Item{{ShortDescription: "Mountain Dew 12PK", Price: "6.49"}} }
	points = calculator.CalculatePoints(receipt)
	expected = 0
	if points != expected {
		t.Errorf("expected %d points, got %d", expected, points)
	}

  }




