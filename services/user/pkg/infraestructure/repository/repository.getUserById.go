package repository

import (
	"context"
	"user/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r mongoRepository) GetUserById(userId interface{}) (models.UserPublic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("Users")
	userObjecId, err := primitive.ObjectIDFromHex(userId.(string))
	if err != nil {
		return models.UserPublic{}, err
	}
	var user models.UserPublic
	filter := bson.M{"_id": userObjecId}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return models.UserPublic{}, ErrUserNotFound
	}
	return user, nil

}
