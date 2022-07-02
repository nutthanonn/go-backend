package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) GetProduct(context *gin.Context) {
	var product []models.Products

	if result := h.DB.Table("products").Select("product_name").Scan(&product); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": result.Error,
		})
		return
	}

	context.JSON(http.StatusOK, &product)
}
