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

	router.GET("product/:id", h.GetProduct)
	router.POST("product_type", h.CreateProductType)
	router.POST("product", h.CreateProduct)
	router.PUT("product/:id", h.UpdateProduct)
	router.DELETE("product/:id", h.DeleteProduct)
	router.Run("localhost:8000")
}
