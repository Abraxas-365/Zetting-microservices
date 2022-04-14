package repository

import (
	"context"
	"time"
	"work-request/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) CreateWorkRequest(newWorkRequest models.WorkRequest) (*models.WorkRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	projectObjectId, err := primitive.ObjectIDFromHex(newWorkRequest.ProjectId.(string))
	if err != nil {
		return nil, err
	}

	workerObjectId, err := primitive.ObjectIDFromHex(newWorkRequest.WorkerId.(string))
	if err != nil {
		return nil, err
	}

	ownerObjectId, err := primitive.ObjectIDFromHex(newWorkRequest.OwnerId.(string))
	if err != nil {
		return nil, err
	}
	check := bson.M{}
	filter := bson.M{"project_id": projectObjectId, "worker_id": workerObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		newWorkRequest.Created = time.Now()
		newWorkRequest.Updated = time.Now()
		newWorkRequest.OwnerId = ownerObjectId
		newWorkRequest.WorkerId = workerObjectId
		newWorkRequest.ProjectId = projectObjectId
		newWorkRequest.Status = "P"
		response, err := collection.InsertOne(ctx, newWorkRequest)
		if err != nil {
			return nil, err
		}
		newWorkRequest.ID = response.InsertedID.(primitive.ObjectID).Hex()
		return &newWorkRequest, nil
	}
	return nil, ErrConflict

}
