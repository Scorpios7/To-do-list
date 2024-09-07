package routes

import (
	"github.com/Scorpios7/To-do-list/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitTodoRoute(app *fiber.App) {
	api := app.Group("/todo")
	api.Post("/create", handlers.CreateTodo)
	api.Get("/get/all", handlers.GetTodos)
	api.Get("/get/:id", handlers.GetTodoById)
	api.Post("/update", handlers.UpdateTodo)
	api.Post("/delete/:id", handlers.DeleteTodo)
}
