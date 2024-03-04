package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) RefreshToken( UUID string, token string) (bool, error) {
	// bool = true - success update token
	// bool = false - failed (not exist)

	filter := bson.M{"UUID": UUID, "token": token}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result := r.collection.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("ErrNoDocuments?", result)
			return false, nil
		}
		log.Println("err?", result)
		return false, err
	}

	log.Println("result?", result)
	return true, nil
}
