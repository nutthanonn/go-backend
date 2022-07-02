package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) CreateProduct(context *gin.Context) {
	var product models.Products

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}

	product.Update_at = time.Now()
	product.Create_at = time.Now()

	if result := h.DB.Create(&product); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	context.JSON(http.StatusCreated, &product)

}
