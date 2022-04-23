package main

import (
	"fmt"
	"os"
	"work-request/internal/rabbit"
	"work-request/pkg/application"
	"work-request/pkg/core/service"
	"work-request/pkg/infraestructure/mqueue/publisher"
	"work-request/pkg/infraestructure/repository"
	"work-request/pkg/infraestructure/rest/handlers"
	"work-request/pkg/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	mongoUri := os.Getenv("MONGODB_URI")
	mqUri := os.Getenv("MQ_URI")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "WorkRequests")
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print("Could not create RabbitMQ")
		os.Exit(1)
	}
	mqpublisher := mqpublisher.NewMQPublisher(mq)
	service := service.NewWorRequesttService(repo)
	application := application.NewWorkRequestApplication(repo, mqpublisher, service)
	handler := handlers.NewWorkRequestHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.WorkRequestRoute(app, handler) //User routes
	fmt.Println("inicando en puerto 3003")

	app.Listen(":3003")
}
