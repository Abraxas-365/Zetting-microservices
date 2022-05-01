package repository

import (
	"context"
	"work-request/workrequest/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mongoRepository) GetWorkRequests(userId uuid.UUID, status string, page int, number int, document string) (models.WorkRequests, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	var workRequests models.WorkRequests
	options := options.Find()
	options.SetLimit(int64(number))
	options.SetSkip((int64(page) - 1) * int64(number))
	filter := bson.D{primitive.E{Key: document, Value: userId}, primitive.E{Key: "status", Value: status}}
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &workRequests); err != nil {
		return nil, err
	}
	return workRequests, nil
}
