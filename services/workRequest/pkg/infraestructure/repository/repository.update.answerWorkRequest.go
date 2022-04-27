package repository

import (
	"context"
	"work-request/pkg/core/events"
	"work-request/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) AnswerWorkRequest(workRequest models.WorkRequest) (events.Event, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	workerObjectId, err := primitive.ObjectIDFromHex(workRequest.Worker.ID.(string))
	if err != nil {
		return nil, err
	}
	workRequestObjectId, err := primitive.ObjectIDFromHex(workRequest.ID.(string))
	if err != nil {
		return nil, err
	}

	updateQuery := bson.M{
		"$set": bson.M{"status": workRequest.Status},
	}
	filter := bson.M{"_id": workRequestObjectId, "worker._id": workerObjectId, "status": "P"}
	_, err = collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return nil, err
	}
	workRequestFromdb := new(models.WorkRequest)
	filter = bson.M{"_id": workRequestObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&workRequestFromdb); err != nil {
		return nil, err
	}

	switch workRequestFromdb.Status {
	case "A":
		return events.WorkrequestAccepted{
			ID:      workRequestFromdb.ID,
			Owner:   workRequestFromdb.Owner,
			Worker:  workRequestFromdb.Worker,
			Project: workRequestFromdb.Project,
		}, nil
	case "B":
		return events.WorkrequestDenied{
			ID:      workRequestFromdb.ID,
			Owner:   workRequestFromdb.Owner,
			Worker:  workRequestFromdb.Worker,
			Project: workRequestFromdb.Project,
		}, nil
	}
	return nil, err
}
