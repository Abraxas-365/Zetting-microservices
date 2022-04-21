package repository

import (
	"context"
	"projects/pkg/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) GetProjectByProjectId(projectId interface{}) (models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	projectObjecId, err := primitive.ObjectIDFromHex(projectId.(string))
	if err != nil {
		return models.Project{}, err
	}
	var project models.Project
	filter := bson.M{"_id": projectObjecId}
	if err := collection.FindOne(ctx, filter).Decode(&project); err != nil {
		return models.Project{}, err
	}
	return project, nil
}
