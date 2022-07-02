package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) DeleteProduct(c *gin.Context) {
	var product models.Products
	var id = c.Param("id")

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}

	if result := h.DB.Where("product_id=?", id).Delete(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &product)
}
