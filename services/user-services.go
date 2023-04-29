package services

import (
	"context"
	"money-tracker/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *models.User, userCollection *mongo.Collection) (*models.User, error) {
	dbCtx := context.Background()

	_, err := userCollection.InsertOne(dbCtx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
