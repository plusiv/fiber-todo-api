package todoHandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plusiv/fiber-todo-api/database"
	"github.com/plusiv/fiber-todo-api/internals/model"
)

// GetTodos godoc
//
//	@Summary		Get all ToDos
//	@Description	Get all ToDos
//	@Accept			json
//	@Produce		json
//	@Router			/todos/ [get]
func GetTodos(c *fiber.Ctx) error {
	db := database.DB
	var todos []model.ToDo

	db.Find(&todos)

	if len(todos) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No ToDos found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "ToDos found", "data": todos})
}

func GetTodo(c *fiber.Ctx) error {
	db := database.DB

	var todo model.ToDo

	todoId := c.Params("todoId")
	db.Find(&todo, "id = ?", todoId)

	if todo.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Todo with Id %s not found", todoId), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "TODOs found", "data": todo})
}
