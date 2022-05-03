package repository

import (
	"context"
	"fmt"
	"notifications/notification/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *mongoRepository) GetNotificationById(notificationId uuid.UUID) (models.LookupNotification, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	lookupNotifications := models.LookupNotifications{}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: notificationId}}}}
	optionsLimit := bson.D{{Key: "$limit", Value: 1}}
	lookupNotified := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "NotificationUsers"}, {Key: "localField", Value: "notified"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "notified"}}}}
	lookupNotifier := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "NotificationUsers"}, {Key: "localField", Value: "notifier"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "notifier"}}}}
	unwindNotified := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$notified"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	unwindNotifier := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$notifier"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{optionsLimit, matchStage, lookupNotified, lookupNotifier, unwindNotified, unwindNotifier})
	if err = cur.All(ctx, &lookupNotifications); err != nil {
		return models.LookupNotification{}, err
	}
	fmt.Println(&lookupNotifications)
	if len(lookupNotifications) <= 0 {
		return models.LookupNotification{}, ErrNotFound
	}

	return *lookupNotifications[0], nil

}
