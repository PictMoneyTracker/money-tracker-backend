package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id          primitive.ObjectID `json:"id"`
	UserId      primitive.ObjectID `json:"userId"`
	Amount      int32              `json:"amount"`
	Category    string             `json:"category"`
	SpendFrom   string             `json:"spendFrom"`
	CreatedAt   time.Time          `json:"createdAt"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
}
