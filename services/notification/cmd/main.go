package main

import (
	"fmt"
	"notifications/pkg/application"
	"notifications/pkg/infraestructure/mqueue"
	"notifications/pkg/infraestructure/repository"
	"notifications/pkg/infraestructure/rest/handlers"
	"notifications/pkg/infraestructure/rest/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	mongoUri := os.Getenv("MONGODB_URI")
	mqUri := os.Getenv("MQ_URI")
	mqChannelName := os.Getenv("MQ_CHANNEL_NAME")

	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Notifications")
	service := service.NewNotificationService(repo)
	handler := handlers.NewNotificationHandler(service)
	mq, err := mqueue.NewMQueue(mqUri, mqChannelName, service)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	go mq.ConsumerWorkRequest()
	app := fiber.New()
	app.Use(logger.New())
	routes.NotificationRoute(app, handler) //User routes
	//Routes.
	fmt.Println("inicando en puerto 3002")
	app.Listen(":3002")
}
