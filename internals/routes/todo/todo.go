package todoRoutes

import (
	"github.com/gofiber/fiber/v2"
	todoHandler "github.com/plusiv/fiber-todo-api/internals/handlers/todo"
)

func SetupTodoRoutes(router fiber.Router){
	todo := router.Group("/todos")
	todo.Get("/", todoHandler.GetTodos)
}