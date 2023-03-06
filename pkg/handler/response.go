package handler

import (
	todo_api "github.com/KuratovIgor/todo-api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type getAllListsResponse struct {
	Data []todo_api.TodoList `json:"data"`
}

type getListByIdResponse struct {
	Data todo_api.TodoList `json:"data"`
}

type getAllItemsResponse struct {
	Data []todo_api.TodoItem `json:"data"`
}

type getItemByIdResponse struct {
	Data todo_api.TodoItem `json:"data"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)

	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
