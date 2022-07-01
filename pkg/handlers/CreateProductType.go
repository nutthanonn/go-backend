package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) CreateProductType(context *gin.Context) {
	var product_type models.Product_type
	product_type.Create_at = time.Now()
	product_type.Update_at = time.Now()

	if err := context.ShouldBindJSON(&product_type); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}

	if result := h.DB.Table("product_types").Create(&product_type); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, &product_type)
}
