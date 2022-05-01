package repository

import (
	"context"
	"projects/project/core/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoRepository) GetProjectByProjectId(projectId uuid.UUID) (models.Project, error) {
	// TODO join with users db
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var project models.Project
	filter := bson.M{"_id": projectId}
	if err := collection.FindOne(ctx, filter).Decode(&project); err != nil {
		return models.Project{}, err
	}

	return project, nil
}
