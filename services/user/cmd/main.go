package main

import (
	// "github.com/arsmn/fiber-swagger/v2"
	"os"
	"user/pkg/application"
	"user/pkg/core/service"
	"user/pkg/infraestructure/repository"
	"user/pkg/infraestructure/rest/handlers"
	"user/pkg/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	mongoUri := os.Getenv("MONGODB_URI")

	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Users")
	service := service.NewUserService(repo)
	application := application.NewUserApplication(repo, service)
	handler := handlers.NewUserHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.UsersRoute(app, handler) //User routes

	app.Listen(":3000")
}
