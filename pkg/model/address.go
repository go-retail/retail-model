package model

// Unexported struct (only accessible inside package `model`)
type address struct {
	House   string `json:"ProductName"`
	Street  string `json:"Street"`
	City    string `json:"City"`
	State   string `json:"State"`
	Country string `json:"Country"`
	ZipCode string `json:"zipcode"`
}
