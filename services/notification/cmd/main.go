package main

import (
	"fmt"
	rabbit "notifications/internal/rabbtit"
	"notifications/pkg/application"
	"notifications/pkg/core/service"
	"notifications/pkg/infraestructure/mqueue/consumer/handlers"
	"notifications/pkg/infraestructure/mqueue/consumer/routes"
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

	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Notifications")
	service := service.NewNotificationService(repo)
	application := application.NewNotificationApplication(repo, service)
	handler := handlers.NewNotificationHandler(application)
	//handler of queue
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	mqhandler := mqHandler.NewMQHandler(application)

	mqconsumer_routes.ConsumerRoutes(mq, mqhandler)
	app := fiber.New()
	app.Use(logger.New())
	routes.NotificationRoute(app, handler) //User routes
	//Routes.
	fmt.Println("inicando en puerto 3002")
	app.Listen(":3002")

}
