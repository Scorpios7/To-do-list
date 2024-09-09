package services

import (
	"fmt"

	database "github.com/Scorpios7/To-do-list/internal"
	"github.com/Scorpios7/To-do-list/internal/models"
)

func CreateTodo(todo models.Todo) error {
	db := database.GetDB()
	if err := db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodos() ([]models.Todo, error) {
	db := database.GetDB()
	var todos []models.Todo
	if err := db.Find(&todos).Error; err != nil {
		fmt.Println("Failed to get todos:", err)
		return nil, err
	}
	return todos, nil
}

func GetTodoById(id int) (*models.Todo, error) {
	db := database.GetDB()
	var todo models.Todo
	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func UpdateTodo(todo models.Todo) error {
	db := database.GetDB()
	var savaTarget models.Todo
	if err := db.First(&savaTarget, todo.Id).Error; err != nil {
		return err
	}
	if savaTarget.Id == 0 {
		return fmt.Errorf("todo not exist")
	}
	if todo.Body != "" {
		savaTarget.Body = todo.Body
	}
	savaTarget.Completed = todo.Completed
	if err := db.Save(&savaTarget).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(id int) error {
	db := database.GetDB()
	todo, err := GetTodoById(id)
	if err != nil {
		return err
	}
	if err := db.Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
