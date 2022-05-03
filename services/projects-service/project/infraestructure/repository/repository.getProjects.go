package repository

import (
	"context"
	"projects/project/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*get projects, document = owner or worker*/
func (r *mongoRepository) GetProjects(userId uuid.UUID, document string, page int) (models.LookupProjects, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	lookupProjects := models.LookupProjects{}
	optionsLimit := bson.D{{Key: "$limit", Value: 20}}
	optionsSkip := bson.D{{Key: "$skip", Value: page - 1}}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "owner", Value: userId}}}}
	lookupWorkers := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "ProjectUsers"}, {Key: "localField", Value: "workers"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "workers"}}}}
	lookupOwner := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "ProjectUsers"}, {Key: "localField", Value: "owner"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "owner"}}}}
	unwindWorkers := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$workers"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	unwindOwner := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$owner"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{optionsLimit, optionsSkip, matchStage,
		lookupOwner, lookupWorkers, unwindOwner, unwindWorkers})
	if err = cur.All(ctx, &lookupProjects); err != nil {
		return nil, err
	}

	return lookupProjects, nil
}
