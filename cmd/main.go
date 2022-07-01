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

	router.POST("product_type", h.CreateProductType)
	router.GET("product_type", h.GetProductType)
	router.POST("product", h.CreateProduct)
	router.Run("localhost:8000")
}
