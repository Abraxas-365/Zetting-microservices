package repository

import (
	"context"
	"projects/pkg/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

/*get projects, document = owner or worker*/
func (r *mongoRepository) GetProjects(userId uuid.UUID, document string, page int) (models.Projects, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	var projects models.Projects
	filter := bson.M{document: userId}
	options := options.Find()
	options.SetLimit(20)
	options.SetSkip((int64(page) - 1) * 20)
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
