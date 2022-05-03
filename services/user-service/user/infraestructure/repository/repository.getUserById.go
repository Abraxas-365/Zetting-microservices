package repository

import (
	"context"
	"user/user/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r mongoRepository) GetUserById(userId uuid.UUID) (models.UserPublic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("Users")
	var user models.UserPublic
	filter := bson.M{"_id": userId}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return models.UserPublic{}, ErrUserNotFound
	}
	return user, nil

}
