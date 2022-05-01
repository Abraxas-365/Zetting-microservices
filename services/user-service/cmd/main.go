package main

import (
	// "github.com/arsmn/fiber-swagger/v2"
	"fmt"
	"os"
	"user/internal/rabbit"
	"user/user/application"
	"user/user/core/service"
	"user/user/infraestructure/mqueue/publisher"
	"user/user/infraestructure/repository"
	"user/user/infraestructure/rest/handlers"
	"user/user/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	mongoUri := os.Getenv("MONGODB_URI")

	mqUri := os.Getenv("MQ_URI")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Users")
	service := service.NewUserService(repo)
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	mqpublisher := mqpublisher.NewMQPublisher(mq)
	application := application.NewUserApplication(repo, service, mqpublisher)
	handler := handlers.NewUserHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.UsersRoute(app, handler) //User routes

	app.Listen(":3000")
}
