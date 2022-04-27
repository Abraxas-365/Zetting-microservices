package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) AddUserToProject(userID interface{}, projectId interface{}, document string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	projectObjectId, err := primitive.ObjectIDFromHex(projectId.(string))
	if err != nil {
		return err
	}

	userToAddObjectId, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		return err
	}
	//the data comes from a work reqeust
	filter := bson.M{"_id": projectObjectId}
	update := bson.M{
		"$push": bson.M{
			document: userToAddObjectId,
		},
	}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}
