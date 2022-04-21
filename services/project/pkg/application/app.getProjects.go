package application

import "projects/pkg/core/models"

func (s *projectApplication) GetProjects(userId interface{}, document string, page int) (models.Projects, error) {
	projects, err := s.projectRepo.GetProjects(userId, document, page)
	if err != nil || projects == nil {
		return []*models.Project{}, nil
	}
	return projects, nil

}
