package handlers

import (
	"github.com/Scorpios7/To-do-list/internal/models"
	"github.com/Scorpios7/To-do-list/internal/services"
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}
	if err := services.CreateTodo(todo); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON("Create success")
}

func GetTodos(c *fiber.Ctx) error {
	todos, err := services.GetTodos()
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetTodoById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	todo, err := services.GetTodoById(id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}
	if err := services.UpdateTodo(todo); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON("Update success")
}

func DeleteTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	if err := services.DeleteTodo(id); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON("Delete success")
}
