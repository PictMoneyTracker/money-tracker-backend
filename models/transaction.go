package models

import "time"

type Transaction struct {
	Id          string    `json:"id"`
	Amount      int32     `json:"amount"`
	Category    string    `json:"category"`
	SpendFrom   string    `json:"spendFrom"`
	CreatedAt   time.Time `json:"createdAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
