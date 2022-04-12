package repository

import (
	"context"
	"user/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) GetUserById(userId interface{}) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	userObjecId, err := primitive.ObjectIDFromHex(userId.(string))
	if err != nil {
		return nil, err
	}
	var user models.User
	filter := bson.M{"_id": userObjecId}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil

}
