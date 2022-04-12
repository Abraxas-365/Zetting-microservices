package repository

import (
	"context"
	"user/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) GetUsersByProject(projectId interface{}, document string) (models.Users, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	projectObjecId, err := primitive.ObjectIDFromHex(projectId.(string))

	var users models.Users
	documentProject := "projects." + document
	filter := bson.M{documentProject: bson.A{projectObjecId}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil

}
