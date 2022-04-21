package repository

import (
	"context"
	"time"
	"work-request/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) CreateWorkRequest(newWorkRequest models.WorkRequest) (models.WorkRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	projectObjectId, err := primitive.ObjectIDFromHex(newWorkRequest.Project.ID.(string))
	if err != nil {
		return models.WorkRequest{}, err
	}

	workerObjectId, err := primitive.ObjectIDFromHex(newWorkRequest.Worker.ID.(string))
	if err != nil {
		return models.WorkRequest{}, err
	}

	ownerObjectId, err := primitive.ObjectIDFromHex(newWorkRequest.Owner.ID.(string))
	if err != nil {
		return models.WorkRequest{}, err
	}
	check := bson.M{}
	filter := bson.M{"project._id": projectObjectId, "worker._id": workerObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		newWorkRequest.Created = time.Now()
		newWorkRequest.Updated = time.Now()
		newWorkRequest.Owner.ID = ownerObjectId
		newWorkRequest.Worker.ID = workerObjectId
		newWorkRequest.Project.ID = projectObjectId
		newWorkRequest.Status = "P"
		response, err := collection.InsertOne(ctx, newWorkRequest)
		if err != nil {
			return models.WorkRequest{}, err
		}
		newWorkRequest.ID = response.InsertedID.(primitive.ObjectID).Hex()
		return newWorkRequest, nil
	}
	return models.WorkRequest{}, ErrConflict

}
