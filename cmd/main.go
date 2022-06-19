package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/db"
	"github.com/nutthanonn/golang-webserver/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := gin.Default()

	router.GET("year", h.GetYear)
	router.POST("year", h.CreateYear)
	router.Run("localhost:8000")
}
