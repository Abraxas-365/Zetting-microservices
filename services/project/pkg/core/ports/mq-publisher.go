package ports

import "projects/pkg/core/models"

type ProjectMQPublisher interface {
	NewProject(project models.Project, exchange string, routingKey string) error
}
