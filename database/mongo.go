package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI string
	Database string
}

type MongoDB struct {
	Client *mongo.Client
	Database *mongo.Database
}

func NewMongoDB(config *MongoConfig) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
        return nil, err
    }

	db := client.Database(config.Database)
	
	return &MongoDB{Client: client, Database: db}, nil
}