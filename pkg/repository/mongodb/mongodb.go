package mongodb

import (
	"context"
	"fmt"
	"jwt-service/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Mongo mongo.Client
}

func InitRepository(cfg *config.ServiceConfig) *Repository {

	return &Repository{
		Mongo: *InitMongodb(cfg),
	}
}

func InitMongodb(cfg *config.ServiceConfig) *mongo.Client {

	var client *mongo.Client

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.Service.JWT.DB.Mongo.URI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
