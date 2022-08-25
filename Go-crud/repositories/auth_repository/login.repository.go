package auth_repository

import (
	"context"
	"go-crud/credentials"
	"go-crud/database"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("login")
var ctx = context.Background()

func SearchUser(username string) (credentials.LoginCredentials, error) {
	var err error
	filter := bson.M{"username": username}

	var result credentials.LoginCredentials
	user := collection.FindOne(ctx, filter).Decode(&result)

	if user != nil {
		return credentials.LoginCredentials{}, err
	}
	return result, nil

}
