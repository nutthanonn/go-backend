package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) GetProduct(c *gin.Context) {
	var product []models.Products
	var id = c.Param("id")

	if id == "all" {
		if result := h.DB.Find(&product); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": result.Error,
			})
			return
		}
	} else {
		if result := h.DB.Where("product_id=?", id).First(&product); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": result.Error,
			})
			return
		}
	}
	c.JSON(http.StatusOK, &product)
}
