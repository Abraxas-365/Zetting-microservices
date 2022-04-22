package repository

import (
	"context"
	"notifications/pkg/core/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) CreateNotification(newNotification models.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	notifierObjectId, err := primitive.ObjectIDFromHex(newNotification.NotifierUser.ID.(string))
	if err != nil {
		return err
	}
	notifiedObjectId, err := primitive.ObjectIDFromHex(newNotification.NotifiedUser.ID.(string))
	if err != nil {
		return err
	}

	workRequestObjectId, err := primitive.ObjectIDFromHex(newNotification.WorkRequest.ID.(string))
	if err != nil {
		return err
	}

	projectObjectId, err := primitive.ObjectIDFromHex(newNotification.WorkRequest.Project.ID.(string))
	if err != nil {
		return err
	}
	newNotification.Created = time.Now()
	newNotification.Updated = time.Now()
	newNotification.NotifierUser.ID = notifierObjectId
	newNotification.NotifiedUser.ID = notifiedObjectId
	newNotification.WorkRequest.ID = workRequestObjectId
	newNotification.WorkRequest.Project.ID = projectObjectId
	newNotification.Read = false
	_, err = collection.InsertOne(ctx, newNotification)
	if err != nil {
		return err
	}

	return nil

}
