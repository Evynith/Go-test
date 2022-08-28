package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "myuser"
	pwd      = "password"
	host     = "localhost"
	port     = 27017
	database = "futbolfifa"
)

func GetCollection(collection string) *mongo.Collection {
	//mongodb://myuser:password@localhost:27017/futbolfifa?authMechanism=DEFAULT
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", usr, pwd, host, port, database)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}
