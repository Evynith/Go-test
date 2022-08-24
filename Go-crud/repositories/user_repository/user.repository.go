package user_repository

import (
	"context"
	"go-crud/database"
	"time"

	m "go-crud/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(User m.User) error {

	var err error
	User.CreatedAt = time.Now()
	User.UpdatedAt = time.Now()
	_, err = collection.InsertOne(ctx, User)
	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Users, error) {
	filter := bson.D{}
	elems, err := collection.Find(ctx, filter)
	var users m.Users
	if err != nil {
		return nil, err
	}

	for elems.Next(ctx) {
		var user m.User
		err = elems.Decode(&user)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func ReadOne(userId string) (m.User, error) {
	var err error
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}

	var result m.User
	user := collection.FindOne(ctx, filter).Decode(&result)

	if user != nil {
		return m.User{}, err
	}
	return result, nil

}

func Update(user m.User, userId string) error {
	var err error
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {
	var err error
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
