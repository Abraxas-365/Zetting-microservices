package repository

import (
	"context"
	"work-request/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) AnswerWorkRequest(workRequest models.WorkRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	workerObjectId, err := primitive.ObjectIDFromHex(workRequest.Worker.(string))
	if err != nil {
		return err
	}
	projectObjectId, err := primitive.ObjectIDFromHex(workRequest.Project.(string))
	if err != nil {
		return err
	}

	ownerObjectId, err := primitive.ObjectIDFromHex(workRequest.Owner.(string))
	if err != nil {
		return err
	}

	check := bson.M{}
	filter := bson.M{"project": projectObjectId, "worker": workerObjectId, "status": workRequest.Status}
	updateQuery := bson.M{
		"$set": bson.M{"status": workRequest.Status},
	}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		filter := bson.M{"project": projectObjectId, "worker": workerObjectId}
		_, err = collection.UpdateOne(ctx, filter, updateQuery)
		if err != nil {
			return err
		}
		/*TODO eliminate this and call from the service to RabbitMQ to sned to project service this action*/
		collection = r.client.Database(r.database).Collection("Projects")
		filter = bson.M{"_id": projectObjectId, "owners": bson.A{ownerObjectId}}

		updateQuery = bson.M{
			"$push": bson.M{
				"workers": workerObjectId,
			},
		}
		if _, err := collection.UpdateOne(ctx, filter, updateQuery); err != nil {
			return err
		}
	}

	return nil
}
