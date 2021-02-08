package handler

import (
	"net/http"

	"github.com/Habibullo-1999/rest_api_gin"
	"github.com/gin-gonic/gin"

)

func (h *Handler) signUp(c *gin.Context) {
	var input rest_api_gin.User

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

func (h *Handler) signIn(ctx *gin.Context) {

}