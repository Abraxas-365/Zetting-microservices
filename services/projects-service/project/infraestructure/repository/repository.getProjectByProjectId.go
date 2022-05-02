package repository

import (
	"context"
	"projects/project/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *mongoRepository) GetProjectByProjectId(projectId uuid.UUID) (models.LookupProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var lookupProjects models.LookupProjects

	lookupOwner := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "ProjectUsers"}, {Key: "localField", Value: "owner"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "owner"}}}}
	unwindOwner := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$owner"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{lookupOwner, unwindOwner})
	if err = cur.All(ctx, &lookupProjects); err != nil {
		return models.LookupProject{}, err
	}
	if len(lookupProjects) > 0 {
		return *lookupProjects[0], nil
	}
	return models.LookupProject{}, nil

}
