package main

import (
	"fmt"
	"os"
	"projects/internal/rabbit"
	projectApp "projects/project/application"
	projectService "projects/project/core/service"
	projectMQPublisher "projects/project/infraestructure/mqueue/publisher"
	projectRepo "projects/project/infraestructure/repository"
	projectRestHandler "projects/project/infraestructure/rest/handlers"
	projectRoutes "projects/project/infraestructure/rest/routes"
	userApp "projects/user/application"
	userMQHandler "projects/user/infraestructure/mqueue/consumer/handlers"
	userMQRoutes "projects/user/infraestructure/mqueue/consumer/routes"
	userRepo "projects/user/infraestructure/repository"
	userRestHandler "projects/user/infraestructure/rest/handlers"
	userRoutes "projects/user/infraestructure/rest/routes"
	workrequestApp "projects/workrequest/application"
	workrequestMQHandler "projects/workrequest/infraestructure/mqueue/consumer/handlers"
	workrequestMQRoutes "projects/workrequest/infraestructure/mqueue/consumer/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	//Infraestructure
	mqUri := os.Getenv("MQ_URI")
	mongoUri := os.Getenv("MONGODB_URI")
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//Repositories
	Projectrepo, _ := projectRepo.NewMongoRepository(mongoUri, "Zetting", 10, "Projects")
	UserRepo, _ := userRepo.NewMongoRepository(mongoUri, "Zetting", 10, "ProjectUsers")

	//Publishers
	ProjectMQPublisher := projectMQPublisher.NewMQPublisher(mq)

	//Services
	ProjectService := projectService.NewProjectService(Projectrepo)

	//Applications
	UserApplication := userApp.NewUserApplication(UserRepo)
	ProjectApplication := projectApp.NewProjectApplication(Projectrepo, ProjectMQPublisher, ProjectService, UserApplication)
	WorkRequestApplication := workrequestApp.NewWorkRequestApplication(ProjectApplication)

	//MQHandlers
	UserMQHandler := userMQHandler.NewMQHandler(UserApplication)
	WorkRequestMQHandler := workrequestMQHandler.NewMQHandler(WorkRequestApplication)

	//MQRoutes
	userMQRoutes.ConsumerRoutes(mq, UserMQHandler)
	workrequestMQRoutes.ConsumerRoutes(mq, WorkRequestMQHandler)

	//Rest
	// RestHandler
	ProjectRestHandlers := projectRestHandler.NewProjectHandler(ProjectApplication)
	UserRestHandlers := userRestHandler.NewUserHandler(UserApplication)
	app := fiber.New()
	app.Use(logger.New())
	projectRoutes.ProjectsRoute(app, ProjectRestHandlers)
	userRoutes.UserRoute(app, UserRestHandlers)

	fmt.Println("inicando en puerto 3000")
	app.Listen(":3001")

}
