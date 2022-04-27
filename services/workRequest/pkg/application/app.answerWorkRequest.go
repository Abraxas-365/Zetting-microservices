package application

import (
	"fmt"
	"os"
	"work-request/pkg/core/models"
)

func (s *workRequestApplication) AnswerWorkRequest(workRequest models.WorkRequest) error {

	//VAlidate the data
	if err := workRequest.Validate(); err != nil {
		return err
	}
	//Safe the data via repository
	workRequest, err := s.projectRepo.AnswerWorkRequest(workRequest)
	if err != nil {
		return err
	}
	//sent to rabbitMQ
	fmt.Println("PRojectId", workRequest.Project.ID)
	workRequest.AnserWorkrequest()
	fmt.Println("Work request answer: ", workRequest)
	if err := s.mqpublisher.AnswerWorkRequest(workRequest, "WorkRequest", "anser_workrequest"); err != nil {
		//TODO: do something with the error
		fmt.Println("Rabbit consumer closed - critical Error")
		os.Exit(1)
	}

	return nil

}
