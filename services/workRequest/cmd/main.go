package main

import (
	"fmt"
	"os"
	"work-request/pkg/application"
	"work-request/pkg/infraestructure/mqueue"
	"work-request/pkg/infraestructure/repository"
	"work-request/pkg/infraestructure/rest/handlers"
	"work-request/pkg/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	mongoUri := os.Getenv("MONGODB_URI")
	mqUri := os.Getenv("MQ_URI")
	mqChannelName := os.Getenv("MQ_CHANNEL_NAME")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "WorkRequests")
	mq, _ := mqueue.NewMQueue(mqUri, mqChannelName)
	service := service.NewWorkRequestService(repo, mq)
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
