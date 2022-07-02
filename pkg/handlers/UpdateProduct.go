package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) UpdateProduct(c *gin.Context) {
	var product models.Products
	var id = c.Param("id")

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}
	product.Update_at = time.Now()

	if result := h.DB.Model(&product).Where("product_id=?", id).Updates(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &product)
}
