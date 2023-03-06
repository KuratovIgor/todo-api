package repository

import (
	"fmt"
	todo_api "github.com/KuratovIgor/todo-api"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todo_api.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err = row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListsItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId int, listId int) ([]todo_api.TodoItem, error) {
	var items []todo_api.TodoItem

	getItemsQuery := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id=$1 AND ul.user_id=$2",
		todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.Select(&items, getItemsQuery, listId, userId)

	return items, err
}

func (r *TodoItemPostgres) GetById(userId int, itemId int) (todo_api.TodoItem, error) {
	var item todo_api.TodoItem

	getItemQuery := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id=$1 AND ul.user_id=$2",
		todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.Get(&item, getItemQuery, itemId, userId)

	return item, err
}

func (r *TodoItemPostgres) Update(userId int, itemId int, item todo_api.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if item.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *item.Title)
		argId++
	}

	if item.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *item.Description)
		argId++
	}

	if item.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *item.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	updateItemQuery := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id=$%d AND ti.id=$%d",
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(updateItemQuery, args...)

	return err
}

func (r *TodoItemPostgres) Delete(userId int, itemId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id=$1 AND ti.id=$2",
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(deleteQuery, userId, itemId)

	return err
}
