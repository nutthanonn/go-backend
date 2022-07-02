package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) CreateProductType(c *gin.Context) {
	var product_type models.Product_types

	if err := c.BindJSON(&product_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}

	product_type.Create_at = time.Now()
	product_type.Update_at = time.Now()

	if result := h.DB.Create(&product_type); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &product_type)
}
