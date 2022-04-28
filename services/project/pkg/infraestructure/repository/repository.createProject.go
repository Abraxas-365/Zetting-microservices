package repository

import (
	"context"
	"fmt"
	"projects/pkg/core/events"
	"projects/pkg/core/models"
)

func (r *mongoRepository) CreateProject(newProject models.Project) (events.Event, error) {
	fmt.Println("--CreateProjectRepo--", newProject.Name)
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	_, err := collection.InsertOne(ctx, newProject)
	if err != nil {
		return nil, err
	}

	return events.ProjectCreated{
		ProjectName: string(newProject.Name),
		Image:       newProject.Image,
		Color:       newProject.Color,
	}, nil

}
