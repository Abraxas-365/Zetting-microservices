package repository

import (
	"context"
	"notifications/pkg/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) GetNotificationById(notificationId uuid.UUID) (models.Notification, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	var notification models.Notification
	filter := bson.M{"_id": notificationId}
	if err := collection.FindOne(ctx, filter).Decode(&notification); err != nil {
		return models.Notification{}, err
	}
	return notification, nil

}
