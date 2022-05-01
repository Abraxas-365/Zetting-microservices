package repository

import (
	// "context"
	"github.com/google/uuid"
	// "go.mongodb.org/mongo-driver/bson"
	"work-request/workrequest/core/events"
)

//TODO: FIXME

func (r *mongoRepository) AnswerWorkRequest(workRequest uuid.UUID, workerId uuid.UUID, status string) (events.Event, error) {

	// ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	// defer cancel()
	// collection := r.client.Database(r.database).Collection(r.collection)

	// updateQuery := bson.M{
	// 	"$set": bson.M{"status": workRequest.Status},
	// }
	// filter := bson.M{"_id": workRequest.ID, "worker._id": workRequest.Worker.ID, "status": "P"}
	// _, err := collection.UpdateOne(ctx, filter, updateQuery)
	// if err != nil {
	// 	return nil, err
	// }

	// switch workRequest.Status {
	// case "A":
	// 	return events.WorkrequestAccepted{
	// 		ID:      workRequest.ID,
	// 		Owner:   workRequest.Owner,
	// 		Worker:  workRequest.Worker,
	// 		Project: workRequest.Project,
	// 	}, nil
	// case "D":
	// 	return events.WorkrequestDenied{
	// 		ID:      workRequest.ID,
	// 		Owner:   workRequest.Owner,
	// 		Worker:  workRequest.Worker,
	// 		Project: workRequest.Project,
	// 	}, nil
	// }
	return nil, nil
}
