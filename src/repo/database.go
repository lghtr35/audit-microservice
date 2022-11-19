package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func InitializeConnection(cs string) *Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cs))
	if err != nil {
		panic(err)
	}
	res := new(Database)
	res.Client = client
	return res
}
