package repository

import (
	"context"
	"notifications/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) IsNotificationExist(newNotification models.Notification) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	check := bson.M{}
	filter := bson.M{"notifier._id": newNotification.NotifierUser.ID, "reference": newNotification.ReferenceId, "notified._id": newNotification.NotifiedUser.ID}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		return false
	}
	return true
}
