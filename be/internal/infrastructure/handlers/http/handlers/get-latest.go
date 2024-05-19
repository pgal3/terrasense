package handlers

import (
	"net/http"
	"strconv"

	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	"github.com/PaoloEG/terrasense/internal/core/services"
	http_mappers "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/http/mappers"
	"github.com/gin-gonic/gin"
)

func GetLatestHandler(service *services.MeasurementsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		chipIDString := c.Param("chipID")
		val, err := strconv.ParseInt(chipIDString, 10, 32)
		if err != nil {
			c.Error(&errors.InvalidInputError{
				Message: "Invalid chipID parameter",
				Details: map[string]any{"original": err.Error()},
			}) 
			return
		}
		chipID := int32(val)
		latestTelemetry, serviceError := service.GetLatestMeasurement(chipID)
		if serviceError != nil {
			c.Error(serviceError) 
			return
		}
		c.JSON(http.StatusOK, http_mappers.ToTelemetryResponse(latestTelemetry))
	}
}
