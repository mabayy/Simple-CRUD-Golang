package main

import (
	"tugas-dua/database"
	"tugas-dua/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.DBInit()
	h := handlers.New(db)
	router := gin.Default()

	router.GET("/orders", h.GetOrder)
	router.POST("/orders", h.CreateOrder)
	router.PUT("/orders/:id", h.UpdateOrder)
	router.DELETE("/orders/:id", h.DeleteOrder)
	router.Run(":3000")

}
