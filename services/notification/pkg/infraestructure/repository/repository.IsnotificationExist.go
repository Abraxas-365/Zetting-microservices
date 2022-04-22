package repository

import (
	"context"
	"notifications/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) IsNotificationExist(newNotification models.Notification) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	notifierObjectId, err := primitive.ObjectIDFromHex(newNotification.NotifierUser.ID.(string))
	if err != nil {
		return true
	}
	notifiedObjectId, err := primitive.ObjectIDFromHex(newNotification.NotifiedUser.ID.(string))
	if err != nil {
		return true
	}

	workRequestObjectId, err := primitive.ObjectIDFromHex(newNotification.WorkRequest.ID.(string))
	if err != nil {
		return true
	}

	check := bson.M{}
	filter := bson.M{"notifier._id": notifierObjectId, "reference": workRequestObjectId, "notified._id": notifiedObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		return false
	}
	return true
}
