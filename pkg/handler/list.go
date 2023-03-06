package handler

import (
	todo_api "github.com/KuratovIgor/todo-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary		Create todo list
// @Security ApiKeyAuth
// @Tags			todo lists
// @Description	Create todo list
// @ID				create-list
// @Accept			json
// @Produce		json
// @Param			input	body todo_api.TodoList	true	"todo list info"
// @Success		200		{integer}	integer		1
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo_api.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})
}

// @Summary		Get all todo lists
// @Security ApiKeyAuth
// @Tags			todo lists
// @Description	Get all todo lists
// @ID				get-lists
// @Accept			json
// @Produce		json
// @Success		200		{object}	getAllListsResponse
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary		Get todo list by id
// @Security ApiKeyAuth
// @Tags			todo lists
// @Description	Get todo list by id
// @ID				get-list
// @Accept			json
// @Produce		json
// @Success		200		{object}	getListByIdResponse
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists/:id [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	list, err := h.services.TodoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getListByIdResponse{
		Data: list,
	})
}

// @Summary		Update todo list
// @Security ApiKeyAuth
// @Tags			todo lists
// @Description	Update todo list
// @ID				update-list
// @Accept			json
// @Produce		json
// @Param			input	body todo_api.UpdateListInput	true	"todo list info"
// @Success		200		{object}	statusResponse
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists/:id [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	var input todo_api.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Update(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary		Delete todo list
// @Security ApiKeyAuth
// @Tags			todo lists
// @Description	Delete todo list
// @ID				delete-list
// @Accept			json
// @Produce		json
// @Success		200		{object}	statusResponse
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists/:id [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	err = h.services.TodoList.Delete(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
