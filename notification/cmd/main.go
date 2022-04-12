package main

import (
	"fmt"
	"notifications/pkg/application"
	"notifications/pkg/infraestructure/repository"
	"notifications/pkg/infraestructure/rest/handlers"
	"notifications/pkg/infraestructure/rest/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	mongoUri := os.Getenv("MONGODB_URI")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Notifications")
	service := service.NewNotificationService(repo)
	handler := handlers.NewNotificationHandler(service)
	app := fiber.New()
	app.Use(logger.New())
	routes.NotificationRoute(app, handler) //User routes
	//Routes.
	fmt.Println("inicando en puerto 3002")
	// app.Get("/swagger/*", swagger.HandlerDefault)
	// app.Get("/swagger/*", swagger.New(swagger.Config{
	// 	URL:         "swagger/doc.json",
	// 	DeepLinking: false,
	// }))

	app.Listen(":3002")
}
