package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) UpdateUser(query interface{}, userId interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	userObjecId, err := primitive.ObjectIDFromHex(userId.(string))
	if err != nil {
		return err
	}

	filter := bson.M{"_id": userObjecId}
	updateQuery := bson.M{
		"$set": query,
	}
	fmt.Println(query)
	_, err = collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return err

	}
	return nil
}
