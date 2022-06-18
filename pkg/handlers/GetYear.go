package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) GetYear(context *gin.Context) {
	var years []models.Year

	if result := h.DB.Find(&years); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": result.Error,
		})
		return
	}

	context.JSON(http.StatusOK, &years)
}
