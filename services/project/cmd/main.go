package main

import (
	"fmt"
	"os"
	"projects/internal/rabbit"
	"projects/pkg/application"
	"projects/pkg/infraestructure/mqueue/publisher"
	"projects/pkg/infraestructure/repository"
	"projects/pkg/infraestructure/rest/handlers"
	"projects/pkg/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	mqUri := os.Getenv("MQ_URI")
	mongoUri := os.Getenv("MONGODB_URI")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Projects")
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	mqpublisher := mqpublisher.NewMQPublisher(mq)
	application := application.NewProjectApplication(repo, mqpublisher)
	handlers := handlers.NewProjectHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.ProjectsRoute(app, handlers) //User routes
	fmt.Println("inicando en puerto 3000")

	app.Listen(":3001")
}
