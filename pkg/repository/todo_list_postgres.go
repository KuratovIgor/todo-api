package repository

import (
	"fmt"
	go_rest_api "github.com/KuratovIgor/go-todo-api"
	"github.com/jmoiron/sqlx"
	"log"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list go_rest_api.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	log.Println("USER ID ", userId)

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]go_rest_api.TodoList, error) {
	var lists []go_rest_api.TodoList

	getListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, getListsQuery, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId int, listId int) (go_rest_api.TodoList, error) {
	var list go_rest_api.TodoList

	getListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND tl.id = $2",
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, getListsQuery, userId, listId)

	return list, err
}
