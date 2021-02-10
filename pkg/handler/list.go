package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getAllList(ctx *gin.Context) {

}
func (h *Handler) getListById(ctx *gin.Context) {

}
func (h *Handler) updateList(ctx *gin.Context) {

}
func (h *Handler) deleteList(ctx *gin.Context) {

}
