package repository

import (
	todo_api "github.com/KuratovIgor/todo-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo_api.User) (int, error)
	GetUser(username, password string) (todo_api.User, error)
}

type TodoList interface {
	Create(userId int, list todo_api.TodoList) (int, error)
	GetAll(userId int) ([]todo_api.TodoList, error)
	GetById(userId int, listId int) (todo_api.TodoList, error)
	Update(userId int, listId int, list todo_api.UpdateListInput) error
	Delete(userId int, listId int) error
}

type TodoItem interface {
	Create(listId int, item todo_api.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]todo_api.TodoItem, error)
	GetById(userId int, itemId int) (todo_api.TodoItem, error)
	Update(userId int, itemId int, item todo_api.UpdateItemInput) error
	Delete(userId int, itemId int) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
