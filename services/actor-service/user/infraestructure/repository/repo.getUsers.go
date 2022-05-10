package repository

import (
	"actor-service/user/core/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) GetUsers() (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	users := models.Users{}
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
