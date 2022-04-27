package application

func (s *projectApplication) AddUserToProject(userID interface{}, projectId interface{}, document string) error {

	if err := s.projectRepo.AddUserToProject(userID, projectId, "workers"); err != nil {

		return err
	}

	return nil

}
