package handlers

import (
	"net/http"

	"github.com/PaoloEG/terrasense/internal/core/services"
	"github.com/gin-gonic/gin"
)

func GetRangeHandler(service *services.MeasurementsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusNotImplemented)
	}
}