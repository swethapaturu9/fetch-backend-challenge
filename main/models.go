package main

type Receipt struct {
	Retailer     string `json:"retailer" validate:"required,retailerRegex"`
	PurchaseDate string `json:"purchaseDate" validate:"required,date"`
	PurchaseTime string `json:"purchaseTime" validate:"required,time24"`
	Items        []Item `json:"items" validate:"required,dive"`
	Total        string `json:"total" validate:"required,priceRegex"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required,shortDescRegex"`
	Price            string `json:"price" validate:"required,priceRegex"`
}

