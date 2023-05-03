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

	// index := mongo.IndexModel{
	// 	Keys:    bson.M{"email": 1},
	// 	Options: options.Index().SetUnique(true),
	// }
	// _, err := userCollection.Indexes().CreateOne(context.Background(), index)
	// if err != nil {
	// 	return nil, err
	// }

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

// todo: update user
func UpdateUser(userId string, user *models.User) (int64, error) {
	dbCtx := context.Background()

	foundUser, err := GetUser(userId)

	if err != nil {
		return 0, err
	}

	// patch user
	if user.Email != "" {
		foundUser.Email = user.Email
	}
	if user.Name != "" {
		foundUser.Name = user.Name
	}
	if user.PhotoUrl != "" {
		foundUser.PhotoUrl = user.PhotoUrl
	}
	if user.StipendTotal != 0 {
		foundUser.StipendTotal = user.StipendTotal
	}
	if user.StockTotal != 0 {
		foundUser.StockTotal = user.StockTotal
	}
	if user.AllowanceTotal != 0 {
		foundUser.AllowanceTotal = user.AllowanceTotal
	}

	res, err := userCollection.UpdateOne(dbCtx, bson.M{"id": foundUser.Id}, bson.M{"$set": foundUser})

	if err != nil {
		return 0, err
	}

	if res.ModifiedCount == 0 {
		return 0, nil
	}

	return res.MatchedCount, nil
}
