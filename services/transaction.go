package services

import (
	"context"
	"money-tracker/configs"
	"money-tracker/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionCollection *mongo.Collection = configs.GetCollection(configs.DB, "transaction")

// to add a transaction, create a transaction, append it's id to the user's transactions array, and add the user's id to the transaction's userId field
func AddTransaction(userId string, transaction *models.Transaction) (*models.Transaction, error) {
	dbCtx := context.Background()

	user, err := GetUser(userId)

	if err != nil {
		return nil, err
	}

	// push transaction model to user model transactions array
	user.Transactions = append(user.Transactions, transaction.Id)

	_, err = userCollection.UpdateOne(dbCtx, bson.M{"id": user.Id}, bson.M{"$set": user})

	if err != nil {
		return nil, err
	}

	// add user id to transaction model
	transaction.UserId = user.Id

	_, err = transactionCollection.InsertOne(dbCtx, transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
