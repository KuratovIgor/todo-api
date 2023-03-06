package service

import (
	todo_api "github.com/KuratovIgor/todo-api"
	"github.com/KuratovIgor/todo-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo_api.TodoList) (int, error)
	GetAll(userId int) ([]todo_api.TodoList, error)
	GetById(userId int, listId int) (todo_api.TodoList, error)
	Update(userId int, listId int, list todo_api.UpdateListInput) error
	Delete(userId int, listId int) error
}

type TodoItem interface {
	Create(userId int, listId int, item todo_api.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]todo_api.TodoItem, error)
	GetById(userId int, itemId int) (todo_api.TodoItem, error)
	Update(userId int, itemId int, item todo_api.UpdateItemInput) error
	Delete(userId int, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		TodoList:      NewTodoListService(repository.TodoList),
		TodoItem:      NewTodoItemService(repository.TodoItem, repository.TodoList),
	}
}
