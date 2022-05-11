package repository

import (
	"actor-service/user/core/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mongoRepository) GetUsers(page int) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	users := models.Users{}
	options := options.Find()
	options.SetLimit(20)
	options.SetSkip(20 * (int64(page) - 1))

	cur, err := collection.Find(ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
