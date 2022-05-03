package repository

import (
	"context"
	"user/user/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) GetUserByEmail(email models.Email) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("Users")
	var user models.User
	filter := bson.M{"contact.email": email}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return models.User{}, ErrUserNotFound
	}
	return user, nil
}
