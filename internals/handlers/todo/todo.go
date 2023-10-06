package todoHandler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/plusiv/fiber-todo-api/database"
	"github.com/plusiv/fiber-todo-api/internals/model"
)

func GetTodos(c *fiber.Ctx) error {
	db := database.DB
	var todos []model.TODO

	db.Find(&todos)
	log.Print("here before db2")

	if len(todos) == 0 {
		log.Print("here")
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No todos found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "TODOs found", "data": todos})
}