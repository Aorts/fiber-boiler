package infrastructure

import (
	"context"
	"log"
	"time"

	"github.com/Aorts/fiber-boiler/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.mongodb.org/mongo-driver/v2/mongo/writeconcern"
)

func NewMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI(config.Get().Mongo.URI).SetReadPreference(readpref.PrimaryPreferred()).SetRetryWrites(true).SetWriteConcern(writeconcern.Majority()))
	if err != nil {
		log.Fatalf("failed to connect mongo: %s\n", err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("failed to ping mongo: %s\n", err.Error())
	}

	return client
}
