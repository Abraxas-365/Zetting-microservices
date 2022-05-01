package repository

import (
	"context"
	"projects/project/core/events"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) AddUserToProject(user uuid.UUID, project uuid.UUID) (events.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	//the data comes from a work reqeust
	filter := bson.M{"_id": project}
	update := bson.M{
		"$push": bson.M{
			"workers": user,
		},
	}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return nil, err
	}

	return nil, nil
}
