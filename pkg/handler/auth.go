package handler

import (
	todo_api "github.com/KuratovIgor/todo-api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary		SignUp
// @Tags			auth
// @Description	create account
// @ID				create-account
// @Accept			json
// @Produce		json
// @Param			input	body		todo_api.User	true	"account info"
// @Success		200		{integer}	integer		1
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input todo_api.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary		SignIn
// @Tags			auth
// @Description	auth user
// @ID				auth-user
// @Accept			json
// @Produce		json
// @Param			input	body SignInInput	true	"account info"
// @Success		200		{string}	string
// @Failure		400,404			{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
