package dto

import "time"

type Transaction struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Amount          float64
	Status          string
	Description     string
	Store           string
	CreatedAt       time.Time
}
