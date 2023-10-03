package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();

	app.Get("/ping", func(c *fiber.Ctx) (err error) {
		err = c.SendString("pong");
		return
	})

	app.Listen(":3000");
}