package repository

import (
	"context"
	"notifications/pkg/core/models"
)

func (r *mongoRepository) CreateNotification(newNotification models.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	_, err := collection.InsertOne(ctx, newNotification)
	if err != nil {
		return err
	}

	return nil

}
