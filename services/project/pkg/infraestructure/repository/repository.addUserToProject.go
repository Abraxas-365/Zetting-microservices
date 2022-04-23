package repository

import (
	"context"
	"projects/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) AddUserToProject(addUserData models.AddUserToProject, document string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	ownerObjectId, err := primitive.ObjectIDFromHex(addUserData.Owner.(string))
	if err != nil {
		return err
	}

	projectObjectId, err := primitive.ObjectIDFromHex(addUserData.Project.(string))
	if err != nil {
		return err
	}

	userToAddObjectId, err := primitive.ObjectIDFromHex(addUserData.UserToAdd.(string))
	if err != nil {
		return err
	}
	//the data comes from a work reqeust
	filter := bson.M{"_id": projectObjectId, "owners": bson.A{ownerObjectId}}
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
