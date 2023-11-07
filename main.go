package main

import (
	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/plusiv/fiber-todo-api/database"
	_ "github.com/plusiv/fiber-todo-api/docs"
	"github.com/plusiv/fiber-todo-api/router"
)

// @title Fiber ToDo API
// @version 1.0
// @description Sample API for creating ToDos tasks.
// @termsOfService http://swagger.io/terms/
// @contact.name Jorge Massih
// @contact.email jorgmassih@gmail.com
// @license.name MIT
// @license.url https://github.com/plusiv/fiber-todo-api/blob/main/LICENSE
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	// Connect with database
	database.ConnectDB()

	router.SetupRouter(app)

	app.Get("/ping", func(c *fiber.Ctx) (err error) {
		err = c.SendString("pong")
		return
	})
	app.Get("/docs/*", swagger.HandlerDefault) // default

	app.Listen(":3000")
}
