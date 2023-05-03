package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id             primitive.ObjectID   `json:"id"`
	Email          string               `json:"email"`
	Name           string               `json:"name"`
	PhotoUrl       string               `json:"photoUrl"`
	StipendTotal   int32                `json:"stipendTotal"`
	StockTotal     int32                `json:"stockTotal"`
	AllowanceTotal int32                `json:"allowanceTotal"`
	Stocks         []Stock              `json:"stock"`
	Transactions   []primitive.ObjectID `json:"transactionIds"`
}
