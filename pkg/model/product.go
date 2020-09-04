package model

import "time"

//Product .. A Struct to hold a Sale Transaction
type Product struct {
	Name          string    `json:"ProductName"`
	Catagory      string    `json:"ProductCatagory"`
	ProductUnit   string    `json:"ProductMeasurementUnit"`
	PricePerUnit  float64   `json:"ProductDescription"`
	PriceCurrency string    `json:"Currency"`
	UpdateDate    time.Time `json:"UpdateDate"`
}
