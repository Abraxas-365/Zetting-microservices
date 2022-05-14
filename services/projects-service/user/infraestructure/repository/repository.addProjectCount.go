package repository

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) AddProjectCount(userId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	updateQuery := bson.M{"$inc": bson.M{"projectCount": 1}}

	_, err := collection.UpdateByID(ctx, userId, updateQuery)
	if err != nil {
		return err
	}
	return nil

}
