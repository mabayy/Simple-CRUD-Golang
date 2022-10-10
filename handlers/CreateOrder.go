package handlers

import (
	"net/http"
	"time"
	"tugas-dua/models"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateOrder(ctx *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)
	order.OrderedAt = time.Now()
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	h.DB.Create(&order)
	result = gin.H{
		"result": order,
	}
	ctx.JSON(http.StatusOK, result)
}
