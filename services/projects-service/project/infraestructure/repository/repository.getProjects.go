package repository

import (
	"context"
	"projects/project/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

/*get projects, document = owner or worker*/
func (r *mongoRepository) GetProjects(userId uuid.UUID, document string, page int) (models.Projects, error) {
	//TODO join with user db
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	var projects models.Projects
	filter := bson.D{{document, userId}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "ProjectUsers"}, {"localField", "workers"}, {"foreignField", "_id"}, {"as", "workers"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$workers"}, {"preserveNullAndEmptyArrays", false}}}}

	showLoadedStructCursor, err := collection.Aggregate(ctx, mongo.Pipeline{filter, lookupStage, unwindStage})
	if err != nil {
		panic(err)
	}
	if err = showLoadedStructCursor.All(ctx, &projects); err != nil {
		panic(err)
	}

	return projects, nil
}
