package main

import (
	"fmt"
	"os"
	"projects/pkg/application"
	"projects/pkg/infraestructure/repository"
	"projects/pkg/infraestructure/rest/handlers"
	"projects/pkg/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	mongoUri := os.Getenv("MONGODB_URI")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Projects")
	service := service.NewProjectService(repo)
	handlers := handlers.NewProjectHandler(service)
	app := fiber.New()
	app.Use(logger.New())
	//Routes.
	//
	routes.ProjectsRoute(app, handlers) //User routes
	fmt.Println("inicando en puerto 3000")
	// app.Get("/swagger/*", swagger.HandlerDefault)
	// app.Get("/swagger/*", swagger.New(swagger.Config{
	// 	URL:         "swagger/doc.json",
	// 	DeepLinking: false,
	// }))

	app.Listen(":3001")
}
