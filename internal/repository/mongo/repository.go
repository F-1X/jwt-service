package mongodb

import (
	"context"
	"jwt-service/internal/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	cfg        config.Mongo
}

func New(cfg config.Mongo) *Repository {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(cfg.URI)
	clientOptions.Auth = &options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("err", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("[+] mongo client connected success")

	database := client.Database(cfg.DB)

	collection := database.Collection(cfg.Collection)

	return &Repository{
		client:     client,
		database:   database,
		collection: collection,
		cfg:        cfg,
	}
}
