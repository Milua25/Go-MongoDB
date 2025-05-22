package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// MongoDB Client and DB name required for connection
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

const (
	DbName   = "hephzibah"
	mongoURI = "mongodb://20.220.197.1:27017/" + DbName
)

var Mg MongoInstance

func Connection() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	db := client.Database(DbName)

	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}
