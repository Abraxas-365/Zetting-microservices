package repository

import (
	"context"
	"work-request/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) AnswerWorkRequest(workRequest models.WorkRequest) (models.WorkRequest, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	workerObjectId, err := primitive.ObjectIDFromHex(workRequest.Worker.ID.(string))
	if err != nil {
		return models.WorkRequest{}, err
	}
	workRequestObjectId, err := primitive.ObjectIDFromHex(workRequest.ID.(string))
	if err != nil {
		return models.WorkRequest{}, err
	}

	updateQuery := bson.M{
		"$set": bson.M{"status": workRequest.Status},
	}
	filter := bson.M{"_id": workRequestObjectId, "worker._id": workerObjectId, "status": "P"}
	_, err = collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return models.WorkRequest{}, err
	}
	workRequestFromdb := new(models.WorkRequest)
	filter = bson.M{"_id": workRequestObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&workRequestFromdb); err != nil {
		return models.WorkRequest{}, err
	}

	return *workRequestFromdb, nil
}
