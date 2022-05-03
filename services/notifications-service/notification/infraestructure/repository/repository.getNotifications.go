package repository

import (
	"fmt"
	"notifications/notification/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

func (r *mongoRepository) GetNotifications(userId uuid.UUID, page int) (models.LookupNotifications, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	lookupNotifications := models.LookupNotifications{}

	optionsLimit := bson.D{{Key: "$limit", Value: 20}}
	optionsSkip := bson.D{{Key: "$skip", Value: page - 1}}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "notified", Value: userId}}}}
	lookupNotified := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "NotificationUsers"}, {Key: "localField", Value: "notified"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "notified"}}}}
	lookUpProject := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "NotificationProjects"}, {Key: "localField", Value: "project"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "project"}}}}
	lookupNotifier := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "NotificationUsers"}, {Key: "localField", Value: "notifier"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "notifier"}}}}
	unwindNotified := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$notified"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	unwindNotifier := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$notifier"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	unwindProject := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$project"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{optionsLimit, optionsSkip, matchStage,
		lookupNotified, lookupNotifier, lookUpProject,
		unwindNotified, unwindNotifier, unwindProject})
	if err = cur.All(ctx, &lookupNotifications); err != nil {
		fmt.Println(&lookupNotifications)
		return nil, err
	}

	return lookupNotifications, nil

}
