package repository

import (
	"context"
	"fmt"
	"projects/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) IsProjectExist(newProject models.Project) bool {
	fmt.Println("--CreateProjectRepo--", newProject.Name)
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	check := bson.M{}
	filter := bson.M{"name": newProject.Name, "owners._id": newProject.Owner}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		return false
	}
	return true
}
