package repository

import (
	"context"
	"time"
	"user/pkg/core/models"
)

func (r *mongoRepository) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	user.Created = time.Now()
	user.Updated = time.Now()
	user.Features.Skills = []string{}
	user.Verified = false
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
