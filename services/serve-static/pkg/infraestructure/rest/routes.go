package rest

import "github.com/gofiber/fiber/v2"

func ServeRoute(app *fiber.App) {
	/*SERVE*/
	static := app.Group("/api/static")
	static.Static("/app_default_images", "./static/app_default_images", fiber.Static{
		Browse: true, // default: false
	})
}
