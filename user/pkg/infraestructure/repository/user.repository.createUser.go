package repository

import (
	"context"
	"time"
	"user/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) CreateUser(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	check := new(models.User)
	filter := bson.M{"contact.email": user.Contact.Email}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		user.Created = time.Now()
		user.Updated = time.Now()
		user.Features.Skills = []string{}
		user.Projects.Owner = []primitive.ObjectID{}
		user.Projects.Worker = []primitive.ObjectID{}
		user.Verified = false
		result, err := collection.InsertOne(ctx, user)
		if err != nil {
			return nil, err
		}
		user.ID = result.InsertedID.(primitive.ObjectID).Hex()
		return &user, nil

	}
	return nil, ErrConflict
}
