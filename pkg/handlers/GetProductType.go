package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) GetProductType(context *gin.Context) {
	var product_type []models.Product_types

	if result := h.DB.Table("product_types").Select("*").Scan(&product_type); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": result.Error,
		})
	}

	context.JSON(http.StatusOK, &product_type)
}
