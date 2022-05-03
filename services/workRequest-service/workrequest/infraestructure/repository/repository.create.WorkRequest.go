package repository

import (
	"context"
	"work-request/workrequest/core/events"
	"work-request/workrequest/core/models"
)

func (r *mongoRepository) CreateWorkRequest(newWorkRequest models.WorkRequest) (events.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	_, err := collection.InsertOne(ctx, newWorkRequest)
	if err != nil {
		return nil, err
	}
	return events.WorkrequestCreated{
		ID:      newWorkRequest.ID,
		Owner:   newWorkRequest.Owner,
		Worker:  newWorkRequest.Worker,
		Project: newWorkRequest.Project,
		Status:  string(newWorkRequest.Status),
	}, nil

}
