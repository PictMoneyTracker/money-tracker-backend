package services

import (
	"context"
	"money-tracker/configs"
	"money-tracker/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user")

func CreateUser(user *models.User) (*models.User, error) {
	dbCtx := context.Background()

	_, err := userCollection.InsertOne(dbCtx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(userId string) (*models.User, error) {
	dbCtx := context.Background()

	objId, _ := primitive.ObjectIDFromHex(userId)
	res := userCollection.FindOne(dbCtx, bson.M{"id": objId})

	var user models.User
	if err := res.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
