package repository

import (
	"context"
	"user/user/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) IsUserExsist(email models.Email) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	check := new(models.User)
	filter := bson.M{"contact.email": email}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		return false
	}
	return true
}
