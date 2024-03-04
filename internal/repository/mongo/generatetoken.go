package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) GenerateToken(token string) error {

	document := bson.M{
		"refresh_token": token,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("insetred")
	return nil

}
