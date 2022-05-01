package main

import (
	"fmt"
	"os"
	"work-request/internal/rabbit"
	projectApp "work-request/project/application"
	projectMQHandler "work-request/project/infraestructure/mqueue/consumer/handlers"
	projectMQRoutes "work-request/project/infraestructure/mqueue/consumer/routes"
	projectRepo "work-request/project/infraestructure/repository"
	userApp "work-request/user/application"
	userMQHandler "work-request/user/infraestructure/mqueue/consumer/handlers"
	userMQRoutes "work-request/user/infraestructure/mqueue/consumer/routes"
	userRepo "work-request/user/infraestructure/repository"
	workrequestApp "work-request/workrequest/application"
	workrequestService "work-request/workrequest/core/service"
	"work-request/workrequest/infraestructure/mqueue/publisher"
	workrequestRepo "work-request/workrequest/infraestructure/repository"
	"work-request/workrequest/infraestructure/rest/handlers"
	"work-request/workrequest/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	//Infraestructure
	mongoUri := os.Getenv("MONGODB_URI")
	mqUri := os.Getenv("MQ_URI")
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print("Could not create RabbitMQ")
		os.Exit(1)
	}

	//repository
	Workrequestrepo, _ := workrequestRepo.NewMongoRepository(mongoUri, "Zetting", 10, "WorkRequests")
	UserRepo, _ := userRepo.NewMongoRepository(mongoUri, "Zetting", 10, "WorkRequestsUsers")
	ProjectRepo, _ := projectRepo.NewMongoRepository(mongoUri, "Zetting", 10, "WorkRequestProject")

	//publisher
	mqpublisher := mqpublisher.NewMQPublisher(mq)

	//service
	WorkRequestService := workrequestService.NewWorRequesttService(Workrequestrepo)

	//application
	WorkrequestApp := workrequestApp.NewWorkRequestApplication(Workrequestrepo, mqpublisher, WorkRequestService)
	UserApp := userApp.NewUserApplication(UserRepo)
	ProjectApp := projectApp.NewProjectApplication(ProjectRepo)

	//MQHandlers
	ProjectMQHandler := projectMQHandler.NewMQHandler(ProjectApp)
	UserMQHandler := userMQHandler.NewMQHandler(UserApp)

	//MQRoutes
	projectMQRoutes.ConsumerRoutes(mq, ProjectMQHandler)
	userMQRoutes.ConsumerRoutes(mq, UserMQHandler)

	//rest
	handler := handlers.NewWorkRequestHandler(WorkrequestApp)
	app := fiber.New()
	app.Use(logger.New())
	routes.WorkRequestRoute(app, handler) //User routes
	fmt.Println("inicando en puerto 3003")

	app.Listen(":3003")
}
