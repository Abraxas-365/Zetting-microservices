package repository

import (
	"context"
	"notifications/pkg/core/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) CreateNotification(newNotification models.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	notifierObjectId, err := primitive.ObjectIDFromHex(newNotification.NotifierUserId.(string))
	if err != nil {
		return err
	}
	notifiedObjectId, err := primitive.ObjectIDFromHex(newNotification.NotifiedUserId.(string))
	if err != nil {
		return err
	}

	referenceObjectId, err := primitive.ObjectIDFromHex(newNotification.ReferenceId.(string))
	if err != nil {
		return err
	}

	check := bson.M{}
	filter := bson.M{"notifier_id": notifierObjectId, "reference_id": referenceObjectId, "notified_id": notifiedObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		newNotification.Created = time.Now()
		newNotification.Updated = time.Now()
		newNotification.NotifierUserId = notifierObjectId
		newNotification.NotifiedUserId = notifiedObjectId
		newNotification.ReferenceId = referenceObjectId
		newNotification.Read = false
		_, err := collection.InsertOne(ctx, newNotification)
		if err != nil {
			return err
		}
	}
	return nil

}
