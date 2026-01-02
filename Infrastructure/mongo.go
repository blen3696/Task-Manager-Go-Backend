package infrastructure

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		return nil, err
	}

	// confirm connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to MongoDB")

	return client.Database("taskmanager"), nil
}
