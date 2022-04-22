package repository

import (
	"context"
	"user/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) IsUserExsist(user models.User) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	check := new(models.User)
	filter := bson.M{"contact.email": user.Contact.Email}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		return false
	}
	return true
}
