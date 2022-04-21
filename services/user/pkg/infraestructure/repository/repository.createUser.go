package repository

import (
	"context"
	"time"
	"user/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	check := new(models.User)
	filter := bson.M{"contact.email": user.Contact.Email}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
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
	return ErrConflict
}
