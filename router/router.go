package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	todoRoutes "github.com/plusiv/fiber-todo-api/internals/routes/todo"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())
	apiV1 := api.Group("/v1", logger.New())

	todoRoutes.SetupTodoRoutes(apiV1)

}
