package service

import (
	go_rest_api "github.com/KuratovIgor/go-todo-api"
	"github.com/KuratovIgor/go-todo-api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list go_rest_api.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]go_rest_api.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, listId int) (go_rest_api.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
