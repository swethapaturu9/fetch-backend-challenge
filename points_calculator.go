package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)


type RetailerNamePointsCalculator struct{}

type PointsCalculator interface {
	CalculatePoints(receipt Receipt) int
}

/*
-One point for every alphanumeric character in the retailer name.
-50 points if the total is a round dollar amount with no cents.
-25 points if the total is a multiple of 0.25.
-5 points for every two items on the receipt.
-If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
-6 points if the day in the purchase date is odd.
-10 points if the time of purchase is after 2:00pm and before 4:00pm.
*/


func (r *RetailerNamePointsCalculator) CalculatePoints(receipt Receipt) int {
	points := 0

	for _, char := range receipt.Retailer {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			points++
		}
	}

	return points
}

type RoundDollarPointsCalculator struct{}

func (r *RoundDollarPointsCalculator) CalculatePoints(receipt Receipt) int {
	total := receipt.Total
	total = strings.TrimSpace(total)

	if strings.HasSuffix(total, ".00") {
		return 50
	}

	return 0
}

type QuarterPointsCalculator struct{}

func (q *QuarterPointsCalculator) CalculatePoints(receipt Receipt) int {
	total := receipt.Total
	total = strings.TrimSpace(total)

	var totalFloat float64

	if _, err := fmt.Sscanf(total, "%f", &totalFloat); err != nil {
		return 0
	}

	if math.Mod(totalFloat, 0.25) == 0 {
		return 25
	}

	return 0
}

type ItemsPointsCalculator struct{}

func (i *ItemsPointsCalculator) CalculatePoints(receipt Receipt) int {
	return len(receipt.Items) / 2 * 5
}

type ItemDescriptionPointsCalculator struct{}

func (i *ItemDescriptionPointsCalculator) CalculatePoints(receipt Receipt) int {
	points := 0

	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			var price float64
			fmt.Sscanf(item.Price, "%f", &price)
			points += int(math.Ceil(price * 0.2))
		}
	}

	return points
}

type OddDayPointsCalculator struct{}

func (o *OddDayPointsCalculator) CalculatePoints(receipt Receipt) int {
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)

	if date.Day()%2 != 0 {
		return 6
	}
	return 0
}

type PurchaseTimePointsCalculator struct{}

func (p *PurchaseTimePointsCalculator) CalculatePoints(receipt Receipt) int {
	time, _ := time.Parse("15:04", receipt.PurchaseTime)

	if time.Hour() >= 14 && time.Hour() < 16 {
		return 10
	}
	return 0
}

func CalculatePoints(receipt Receipt) int {
	calculators := []PointsCalculator{
		&RetailerNamePointsCalculator{},
		&RoundDollarPointsCalculator{},
		&QuarterPointsCalculator{},
		&ItemsPointsCalculator{},
		&ItemDescriptionPointsCalculator{},
		&OddDayPointsCalculator{},
		&PurchaseTimePointsCalculator{},
	}

	points := 0

	for _, calculator := range calculators {
		points += calculator.CalculatePoints(receipt)
	}

	return points
}
