package repository

import (
	"context"
	"work-request/workrequest/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) IsWorkrequestExist(newWorkRequest models.WorkRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	check := bson.M{}
	filter := bson.M{"project._id": newWorkRequest.Project.ID, "worker._id": newWorkRequest.Worker.ID}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		return false
	}
	return true
}
