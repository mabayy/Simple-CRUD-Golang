package handlers

import (
	"net/http"
	"tugas-dua/models"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	var (
		order    models.Order
		newOrder models.Order
		result   gin.H
	)

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err := h.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = h.DB.Model(&order).Updates(newOrder).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "update success",
		}
	}
	ctx.JSON(http.StatusOK, result)
}
