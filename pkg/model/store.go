package model

import "time"

//Store .. A Struct to hold a Sale Transaction
type Store struct {
	StoreID    string    `json:"StoreID"`
	Address    address   `json:"address"`
	UpdateDate time.Time `json:"txnDate"`
}
