package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) CreateProduct(context *gin.Context) {
	var product models.Product
	product.Create_at = time.Now()
	product.Update_at = time.Now()

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}

	if result := h.DB.Table("products").Create(&product); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, &product)

}
