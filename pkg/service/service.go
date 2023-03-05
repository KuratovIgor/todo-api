package service

import (
	go_rest_api "github.com/KuratovIgor/go-todo-api"
	"github.com/KuratovIgor/go-todo-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user go_rest_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list go_rest_api.TodoList) (int, error)
	GetAll(userId int) ([]go_rest_api.TodoList, error)
	GetById(userId int, listId int) (go_rest_api.TodoList, error)
}

type TodoItem interface {
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
	}
}
