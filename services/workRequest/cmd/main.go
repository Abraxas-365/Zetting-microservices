package main

import (
	"fmt"
	"os"
	"work-request/internal/rabbit"
	"work-request/pkg/application"
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
		os.Exit(1)
	}
	mqpublisher := mqpublisher.NewMQPublisher(mq)
	service := service.NewWorkRequestService(repo, mqpublisher)
	handler := handlers.NewWorkRequestHandler(service)
	app := fiber.New()
	app.Use(logger.New())
	routes.WorkRequestRoute(app, handler) //User routes
	//Routes.
	fmt.Println("inicando en puerto 3003")
	// app.Get("/swagger/*", swagger.HandlerDefault)
	// app.Get("/swagger/*", swagger.New(swagger.Config{
	// 	URL:         "swagger/doc.json",
	// 	DeepLinking: false,
	// }))

	app.Listen(":3003")
}
