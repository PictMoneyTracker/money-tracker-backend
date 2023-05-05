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
	res := userCollection.FindOne(dbCtx, bson.M{"id": user.Id, "stocks.id": stock.Id})

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

func GetStocks(userId string) ([]models.Stock, error) {

	user, err := GetUser(userId)

	if err != nil {
		return nil, err
	}

	return user.Stocks, nil
}

func DeleteStock(userId string, stockId int32) (int32, error) {
	dbCtx := context.Background()

	user, err := GetUser(userId)

	if err != nil {
		return 0, err
	}

	// check if stock exists on user model stocks array
	res := userCollection.FindOne(dbCtx, bson.M{"stocks.id": stockId})

	if res.Err() != nil {
		return 0, errors.New("Stock does not exist")
	}

	// remove stock from user model stocks array
	// var stock models.Stock

	// for i, s := range user.Stocks {
	// 	if s.Id == stockId {
	// 		stock = s
	// 		user.Stocks = append(user.Stocks[:i], user.Stocks[i+1:]...)
	// 		break
	// 	}
	// }

	_, err = userCollection.UpdateOne(dbCtx, bson.M{"id": user.Id}, bson.M{"$pull": bson.M{"stocks": bson.M{"id": stockId}}})

	if err != nil {
		return 0, err
	}

	return stockId, nil
}
