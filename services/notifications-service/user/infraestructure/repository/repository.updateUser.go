package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) UpdateUser(query interface{}, userId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	filter := bson.M{"_id": userId}
	updateQuery := bson.M{
		"$set": query,
	}
	fmt.Println(query)
	_, err := collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return err

	}
	return nil
}
