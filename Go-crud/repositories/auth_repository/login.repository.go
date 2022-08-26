package auth_repository

import (
	"context"
	"go-crud/database"
	m "go-crud/models"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("login")
var ctx = context.Background()

func SearchUser(username string) (m.Client, error) {
	var err error
	filter := bson.M{"username": username}

	var result m.Client
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return m.Client{}, err
	}
	return result, nil

}

func PersistToken(username string, token string) error {
	var err error
	filter := bson.M{"username": username}
	newT := bson.M{
		"$set": bson.M{"token": token},
	}

	_, err = collection.UpdateOne(ctx, filter, newT)
	if err != nil {
		return err
	}
	return nil
}

func ExistsToken(token string) bool {
	var err error
	filter := bson.M{"token": token}

	var result m.Client
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return false
	}
	return true
}
