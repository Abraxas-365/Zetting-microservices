package repository

import (
	"context"
	"work-request/pkg/core/events"
	"work-request/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) AnswerWorkRequest(workRequest models.WorkRequest) (events.Event, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	updateQuery := bson.M{
		"$set": bson.M{"status": workRequest.Status},
	}
	filter := bson.M{"_id": workRequest.ID, "worker._id": workRequest.Worker.ID, "status": "P"}
	_, err := collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return nil, err
	}

	switch workRequest.Status {
	case "A":
		return events.WorkrequestAccepted{
			ID:      workRequest.ID,
			Owner:   workRequest.Owner,
			Worker:  workRequest.Worker,
			Project: workRequest.Project,
		}, nil
	case "B":
		return events.WorkrequestDenied{
			ID:      workRequest.ID,
			Owner:   workRequest.Owner,
			Worker:  workRequest.Worker,
			Project: workRequest.Project,
		}, nil
	}
	return nil, err
}
