package repository

import (
	"context"
	"fmt"
	"projects/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) IsProjectExist(newProject models.Project, userId interface{}) bool {
	fmt.Println("--CreateProjectRepo--", newProject.Name, userId)
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	userIdObjectId, err := primitive.ObjectIDFromHex(userId.(string))
	if err != nil {
		return true
	}
	check := bson.M{}
	filter := bson.M{"name": newProject.Name, "owners._id": userIdObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		return false
	}
	return true
}
