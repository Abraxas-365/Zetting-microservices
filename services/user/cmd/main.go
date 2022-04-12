package main

import (
	"fmt"
	// "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"user/pkg/application"
	"user/pkg/infraestructure/repository"
	"user/pkg/infraestructure/rest/handlers"
	"user/pkg/infraestructure/rest/routes"
)

func main() {
	mongoUri := os.Getenv("MONGODB_URI")

	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Users")
	service := service.NewUsertService(repo)
	handler := handlers.NewUserHandler(service)
	app := fiber.New()
	app.Use(logger.New())
	routes.UsersRoute(app, handler) //User routes
	//Routes.
	fmt.Println("inicando en puerto 3000")
	// app.Get("/swagger/*", swagger.HandlerDefault)
	// app.Get("/swagger/*", swagger.New(swagger.Config{
	// 	URL:         "swagger/doc.json",
	// 	DeepLinking: false,
	// }))

	app.Listen(":3000")
}
