package repository

import (
	"context"
	"projects/project/core/events"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) AddWorkRequestToProject(projectId uuid.UUID, workRequestId uuid.UUID) (events.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"_id": projectId}
	update := bson.M{
		"$push": bson.M{
			"workrequests": workRequestId,
		},
	}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return nil, err
	}

	return nil, nil
}
