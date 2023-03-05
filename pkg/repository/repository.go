package repository

import (
	go_rest_api "github.com/KuratovIgor/go-todo-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user go_rest_api.User) (int, error)
	GetUser(username, password string) (go_rest_api.User, error)
}

type TodoList interface {
	Create(userId int, list go_rest_api.TodoList) (int, error)
	GetAll(userId int) ([]go_rest_api.TodoList, error)
	GetById(userId int, listId int) (go_rest_api.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
