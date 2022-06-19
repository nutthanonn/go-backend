package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/golang-webserver/pkg/models"
)

func (h *handlers) CreateYear(context *gin.Context) {
	var year models.Year

	if err := context.ShouldBindJSON(&year); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}

	if result := h.DB.Table("year").Create(&year); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, &year)
}
