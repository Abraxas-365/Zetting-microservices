package repository

import (
	"context"
	"fmt"
	"projects/pkg/core/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) CreateProject(newProject *models.Project, userId interface{}) (interface{}, error) {
	fmt.Println("--CreateProjectRepo--", newProject.Name, userId)
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	userIdObjectId, err := primitive.ObjectIDFromHex(userId.(string))
	if err != nil {
		return nil, err
	}
	check := bson.M{}
	filter := bson.M{"name": newProject.Name, "owners": bson.A{userIdObjectId}}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		newProject.Created = time.Now()
		newProject.Updated = time.Now()
		newProject.Owners = []primitive.ObjectID{userIdObjectId}
		newProject.Workers = []primitive.ObjectID{}
		result, err := collection.InsertOne(ctx, newProject)
		if err != nil {
			return nil, err
		}

		id := result.InsertedID.(primitive.ObjectID).Hex()
		return id, nil
	}
	return nil, ErrConflict
}
