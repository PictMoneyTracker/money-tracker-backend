package services

import (
	"context"
	"money-tracker/configs"
	"money-tracker/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetTransactions(userId string) ([]models.Transaction, error) {

	priUserId, err := primitive.ObjectIDFromHex(userId)
	var transactions []models.Transaction

	if err != nil {
		return nil, err
	}

	// get all transactions that have the user's id
	cursor, err := transactionCollection.Find(context.Background(), bson.M{"userid": priUserId})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func DeleteTransaction(userId string, transactionId string) (int64, error) {
	dbCtx := context.Background()

	priUserId, err := primitive.ObjectIDFromHex(userId)
	priTransactionId, err := primitive.ObjectIDFromHex(transactionId)

	if err != nil {
		return 0, err
	}
	// check if user exists
	_, err = GetUser(userId)

	if err != nil {
		return 0, err
	}

	// check if transaction exists on user model transactions array
	found := userCollection.FindOne(dbCtx, bson.M{"transactions": priTransactionId})

	if found.Err() != nil {
		return 0, found.Err()
	}

	// delete transaction from transaction collection
	res, err := transactionCollection.DeleteOne(dbCtx, bson.M{"id": priTransactionId})

	if err != nil {
		return 0, err
	}

	// delete transaction id from user model transactions array
	_, err = userCollection.UpdateOne(dbCtx, bson.M{"id": priUserId}, bson.M{"$pull": bson.M{"transactions": priTransactionId}})

	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

// TODO: patch transaction
func UpdateTransaction(userId string, transactionId string, transaction *models.Transaction) (*models.Transaction, error) {
	dbCtx := context.Background()

	priUserId, err := primitive.ObjectIDFromHex(userId)
	priTransactionId, err := primitive.ObjectIDFromHex(transactionId)

	if err != nil {
		return nil, err
	}

	// check if user exists
	_, err = GetUser(userId)

	if err != nil {
		return nil, err
	}

	// check if transaction exists on user model transactions array
	found := userCollection.FindOne(dbCtx, bson.M{"transactions": priTransactionId})

	if found.Err() != nil {
		return nil, found.Err()
	}

	found = transactionCollection.FindOne(dbCtx, bson.M{"id": priTransactionId})

	if found.Err() != nil {
		return nil, found.Err()
	}

	foundTrx := models.Transaction{}

	err = found.Decode(&foundTrx)

	if err != nil {
		return nil, err
	}

	if transaction.Amount != 0 {
		foundTrx.Amount = transaction.Amount
	}
	if transaction.Category != "" {
		foundTrx.Category = transaction.Category
	}
	if transaction.SpendFrom != "" {
		foundTrx.SpendFrom = transaction.SpendFrom
	}
	if transaction.Title != "" {
		foundTrx.Title = transaction.Title
	}
	if transaction.Description != "" {
		foundTrx.Description = transaction.Description
	}

	// update transaction in transaction collection
	_, err = transactionCollection.UpdateOne(dbCtx, bson.M{"id": priTransactionId}, bson.M{"$set": foundTrx})

	if err != nil {
		return nil, err
	}

	// update transaction in user model transactions array
	_, err = userCollection.UpdateOne(dbCtx, bson.M{"id": priUserId, "transactions.id": priTransactionId}, bson.M{"$set": foundTrx})

	if err != nil {
		return nil, err
	}

	return &foundTrx, nil
}
