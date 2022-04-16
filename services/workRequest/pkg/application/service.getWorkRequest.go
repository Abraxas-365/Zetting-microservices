package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"work-request/pkg/core/models"
)

func (r *workRequestService) GetWorkRequest(workRequestId interface{}) (*models.WorkRequest, error) {

	workRequest, err := r.projectRepo.GetWorkRequest(workRequestId)
	if err != nil {
		return nil, err
	}
	project, err := getProjectInfo(workRequest.Project.(string))
	if err != nil {
		fmt.Println(err.Error())
		// TODO: do something with the error
	}
	workRequest.Project = project

	return workRequest, nil

}

func getProjectInfo(projectId string) (interface{}, error) {
	project := struct {
		Name        string `json:"name,omitempty"`
		Image       string ` json:"image,omitempty"`
		Description string `json:"description,omitempty"`
	}{}

	//go func to call the microservices of project,project
	responseProject, err := http.Get("http://zetting-project:3001/api/projects/id=" + projectId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	responseProjectData, err := ioutil.ReadAll(responseProject.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	json.Unmarshal(responseProjectData, &project)
	return project, nil
}
