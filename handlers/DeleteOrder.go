package handlers

import (
	"net/http"
	"tugas-dua/models"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) DeleteOrder(ctx *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)
	id := ctx.Param("id")
	err := h.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = h.DB.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "delete success",
		}
	}
	ctx.JSON(http.StatusOK, result)
}
