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

	notifierObjectId, err := primitive.ObjectIDFromHex(newNotification.NotifierUser.(string))
	if err != nil {
		return err
	}
	notifiedObjectId, err := primitive.ObjectIDFromHex(newNotification.NotifiedUser.(string))
	if err != nil {
		return err
	}

	referenceObjectId, err := primitive.ObjectIDFromHex(newNotification.Reference.(string))
	if err != nil {
		return err
	}

	check := bson.M{}
	filter := bson.M{"notifier": notifierObjectId, "reference": referenceObjectId, "notified": notifiedObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		newNotification.Created = time.Now()
		newNotification.Updated = time.Now()
		newNotification.NotifierUser = notifierObjectId
		newNotification.NotifiedUser = notifiedObjectId
		newNotification.Reference = referenceObjectId
		newNotification.Read = false
		_, err := collection.InsertOne(ctx, newNotification)
		if err != nil {
			return err
		}
	}
	return nil

}
