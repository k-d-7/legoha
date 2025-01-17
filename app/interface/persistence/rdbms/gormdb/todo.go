package gormdb

import (
	"context"

	"github.com/jinzhu/gorm"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

func NewTodoRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateNewTodo(ctx context.Context, t entity.Todo) error {
	todo := newTodo(t)

	err := r.DB.Create(&todo).Error
	if err != nil {
		return err
	}

	return nil
}

// get all todo
func (r *Repository) ListTodos() ([]entity.Todo, error) {
	var todos []entity.Todo
	err := r.DB.Find(&todos).Error
	if err != nil {
		return todos, err
	}

	result := make([]entity.Todo, 0, len(todos))
	for _, todo := range todos {
		result = append(result, newTodo(todo))
	}

	return result, nil
}

// get todo by id
func (r *Repository) GetTodoByID(id int) (entity.Todo, error) {
	var todo entity.Todo

	err := r.DB.First(&todo, id).Error
	if err != nil {
		return todo, err
	}

	return newTodo(todo), nil
}

// update todo information
func (r *Repository) PatchTodo(id int, t entity.Todo) error {
	var todo entity.Todo
	err := r.DB.First(&todo, id).Error
	if err != nil {
		return err
	}
	err = r.DB.Model(&todo).Updates(t).Error
	r.DB.Save(&todo)
	if err != nil {
		return err
	}

	return nil
}

// delete user by id
func (r *Repository) DeleteTodo(todo entity.Todo) error {

	err := r.DB.First(&todo, todo.ID).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&todo).Error
	if err != nil {
		return err
	}

	return nil
}
