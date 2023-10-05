package main

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/plusiv/fiber-todo-api/database"
)

func main() {
	app := fiber.New();
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}

	// Connect with database
	database.ConnectDB()

	app.Use(swagger.New(cfg))

	app.Get("/ping", func(c *fiber.Ctx) (err error) {
		err = c.SendString("pong");
		return
	})

	app.Listen(":3000");
}