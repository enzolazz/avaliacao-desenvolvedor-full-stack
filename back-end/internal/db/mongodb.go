package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

func ConnectMongoDB(uri, dbName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	Client = client
	Database = client.Database(dbName)

	log.Println("Connected to MongoDB")
}

func DisconnectMongoDB() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := Client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		} else {
			log.Println("Disconnected from MongoDB")
		}
	}
}
