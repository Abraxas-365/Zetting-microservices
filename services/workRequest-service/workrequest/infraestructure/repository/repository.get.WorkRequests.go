package repository

import (
	"context"
	"work-request/workrequest/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func (r *mongoRepository) GetWorkRequests(userId uuid.UUID, status string, page int, number int, document string) (models.LookUpWorkRequests, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	lookUpWorkRequest := models.LookUpWorkRequests{}
	optionsLimit := bson.D{{Key: "$limit", Value: 20}}
	optionsSkip := bson.D{{Key: "$skip", Value: page - 1}}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: document, Value: userId}, {Key: "status", Value: status}}}}
	lookupOwner := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "WorkRequestsUsers"}, {Key: "localField", Value: "owner"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "owner"}}}}
	lookupWorker := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "WorkRequestsUsers"}, {Key: "localField", Value: "worker"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "worker"}}}}
	lookUpProject := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "WorkRequestProject"}, {Key: "localField", Value: "project"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "project"}}}}
	unwindOwner := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$owner"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	unwindWorker := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$worker"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	unwindProject := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$project"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{optionsLimit, optionsSkip,
		matchStage, lookupOwner, lookupWorker, lookUpProject, unwindOwner, unwindWorker, unwindProject})
	if err = cur.All(ctx, &lookUpWorkRequest); err != nil {
		return nil, err
	}
	return lookUpWorkRequest, nil
}
