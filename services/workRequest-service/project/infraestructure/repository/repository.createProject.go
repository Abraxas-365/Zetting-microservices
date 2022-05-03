package repository

import (
	"context"
	"work-request/project/core/models"
)

func (r *mongoRepository) CreateProject(project models.Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, project)
	if err != nil {
		return err
	}

	return nil
}
