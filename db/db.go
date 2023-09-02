package db

import (
	"context"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var ctx = context.Background()

// InitDB initializes the MongoDB client and returns it
func InitDB() (*mongo.Client, error) {
	conf, err := config.LoadConfig("./")
	if err != nil {
		return nil, err
	}
	connectionString := conf.MongoURI()

	// Client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check if the client can ping the server
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Println("DB successfully connected...")
	return client, nil
}
