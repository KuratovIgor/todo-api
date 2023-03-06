package handler

import (
	todo_api "github.com/KuratovIgor/todo-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary		Create todo item
// @Security ApiKeyAuth
// @Tags			todo items
// @Description	Create todo item
// @ID				create-item
// @Accept			json
// @Produce		json
// @Param			input	body todo_api.TodoItem	true	"todo item info"
// @Success		200		{integer}	integer		1
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/:list_id/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	var input todo_api.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": itemId,
	})
}

// @Summary		Get all todo items
// @Security ApiKeyAuth
// @Tags			todo items
// @Description	Get all todo items
// @ID				get-items
// @Accept			json
// @Produce		json
// @Success		200		{object}	getAllItemsResponse
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/:list_id/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

// @Summary		Get todo item by id
// @Security ApiKeyAuth
// @Tags			todo items
// @Description	Get todo item by id
// @ID				get-item
// @Accept			json
// @Produce		json
// @Success		200		{object}	getItemByIdResponse
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/items/:id [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getItemByIdResponse{
		Data: item,
	})
}

// @Summary		Update todo item
// @Security ApiKeyAuth
// @Tags			todo items
// @Description	Update todo item
// @ID				update-item
// @Accept			json
// @Produce		json
// @Param			input	body todo_api.UpdateItemInput	true	"todo item info"
// @Success		200		{object}	statusResponse
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/items/:id [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
	}

	var input todo_api.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.Update(userId, itemId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary		Delete todo item
// @Security ApiKeyAuth
// @Tags			todo items
// @Description	Delete todo item
// @ID				delete-item
// @Accept			json
// @Produce		json
// @Success		200		{object}	statusResponse
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/items/:id [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
