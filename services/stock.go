package services

import (
	"context"
	"errors"
	"money-tracker/models"

	"go.mongodb.org/mongo-driver/bson"
)

func AddStock(userId string, stock *models.Stock) (*models.Stock, error) {
	dbCtx := context.Background()

	user, err := GetUser(userId)

	if err != nil {
		return nil, err
	}
	// check if stock already exists on user model stocks array
	res := userCollection.FindOne(dbCtx, bson.M{"stocks.id": stock.Id})

	if res.Err() == nil {
		return nil, errors.New("Stock already exists")
	}

	// push stock model to user model stocks array
	user.Stocks = append(user.Stocks, *stock)

	_, err = userCollection.UpdateOne(dbCtx, bson.M{"id": user.Id}, bson.M{"$set": user})

	if err != nil {
		return nil, err
	}

	return stock, nil
}
