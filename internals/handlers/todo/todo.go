package todoHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plusiv/fiber-todo-api/database"
	"github.com/plusiv/fiber-todo-api/internals/model"
)

func GetTodos(c *fiber.Ctx) error {
	db := database.DB
	var todos []model.ToDo

	db.Find(&todos)

	if len(todos) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No ToDos found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "TODOs found", "data": todos})
}
