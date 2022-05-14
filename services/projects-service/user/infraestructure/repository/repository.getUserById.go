package repository

import (
	"context"
	"projects/user/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r mongoRepository) GetUserById(userId uuid.UUID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var user models.User
	filter := bson.M{"_id": userId}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return models.User{}, ErrUserNotFound
	}
	return user, nil

}
