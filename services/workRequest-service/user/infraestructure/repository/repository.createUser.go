package repository

import (
	"context"
	"work-request/user/core/models"
)

func (r *mongoRepository) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
