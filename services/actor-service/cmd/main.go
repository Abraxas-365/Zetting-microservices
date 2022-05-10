package main

import (
	"actor-service/internal/rabbit"
	"actor-service/user/application"
	"actor-service/user/infraestructure/mqconsumer/handlers"
	"actor-service/user/infraestructure/mqconsumer/routes"
	"actor-service/user/infraestructure/repository"
	"actor-service/user/infraestructure/rest/handels"
	"actor-service/user/infraestructure/rest/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	mongoUri := os.Getenv("MONGODB_URI")
	mqUri := os.Getenv("MQ_URI")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "ActorUsers")
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	application := application.NewUserApplication(repo)

	//MQHandlers
	UserMQHandler := mqHandler.NewMQHandler(application)

	//MQRoutes
	mqRoutes.ConsumerRoutes(mq, UserMQHandler)
	handler := handlers.NewUserHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.UsersRoute(app, handler) //User routes

	app.Listen(":5001")
}
