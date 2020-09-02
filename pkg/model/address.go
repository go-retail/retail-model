package model

// Unexported struct (only accessible inside package `model`)
type address struct {
	house   string `json:"ProductName"`
	street  string `json:"ProductName"`
	city    string `json:"ProductName"`
	state   string `json:"ProductName"`
	country string `json:"ProductName"`
	zipCode string `json:"zipcode"`
}
