package repository

import (
	"context"
	"work-request/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) IsWorkrequestExist(newWorkRequest models.WorkRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	projectObjectId, err := primitive.ObjectIDFromHex(newWorkRequest.Project.ID.(string))
	if err != nil {
		return true
	}

	workerObjectId, err := primitive.ObjectIDFromHex(newWorkRequest.Worker.ID.(string))
	if err != nil {
		return true
	}

	check := bson.M{}
	filter := bson.M{"project._id": projectObjectId, "worker._id": workerObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		return false
	}
	return true
}
