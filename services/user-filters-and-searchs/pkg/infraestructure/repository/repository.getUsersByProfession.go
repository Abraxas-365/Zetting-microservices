package repository

import (
	"context"
	"user/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mongoRepository) GetUsersByProfession(profession string, page int) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var users models.Users
	userObjecId, err := primitive.ObjectIDFromHex("626649d17c27f186febe8ff0")
	options := options.Find()
	options.SetLimit(20)
	options.SetSkip((int64(page) - 1) * 20)
	// filter := bson.M{"profession.name": profession}
	filter := bson.M{"_id": bson.M{"$ne": userObjecId}, "profession.name": profession}
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil

}
