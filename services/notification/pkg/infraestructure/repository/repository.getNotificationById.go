package repository

import (
	"context"
	"notifications/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) GetNotificationById(notificationId interface{}) (models.Notification, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	notificationObjectId, err := primitive.ObjectIDFromHex(notificationId.(string))
	if err != nil {
		return models.Notification{}, err
	}

	var notification models.Notification
	filter := bson.M{"_id": notificationObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&notification); err != nil {
		return models.Notification{}, err
	}
	return notification, nil

}