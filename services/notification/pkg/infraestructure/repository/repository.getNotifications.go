package repository

import (
	"notifications/pkg/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func (r *mongoRepository) GetNotifications(userId uuid.UUID, page int) (models.Notifications, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	var notifications models.Notifications

	filter := bson.M{"notified._id": userId}
	options := options.Find()
	options.SetLimit(20)
	options.SetSkip((int64(page) - 1) * 20)
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &notifications); err != nil {
		return nil, err
	}

	return notifications, nil

}
