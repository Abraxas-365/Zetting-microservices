package main

import (
	"fmt"
	"static-serve/pkg/infraestructure/rest"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	rest.ServeRoute(app) //User routes
	fmt.Println("inicando en puerto 9000")

	app.Listen(":9000")
}
