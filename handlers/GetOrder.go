package handlers

import (
	"net/http"
	"tugas-dua/models"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetOrder(ctx *gin.Context) {
	var (
		order  []models.Order
		result gin.H
	)
	h.DB.Find(&order)
	if len(order) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": order,
			"count":  len(order),
		}
	}
	ctx.JSON(http.StatusOK, result)
}
