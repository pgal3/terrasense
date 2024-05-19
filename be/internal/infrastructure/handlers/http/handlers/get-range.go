package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	"github.com/PaoloEG/terrasense/internal/core/services"
	http_mappers "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/http/mappers"
	"github.com/gin-gonic/gin"
)

const TIME_LAYOUT = "2006-01-02"

func GetRangeHandler(service *services.MeasurementsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		chipIDString := c.Param("chipID")
		from := c.Query("from")
		to := c.Query("to")

		val, err := strconv.ParseInt(chipIDString, 10, 32)
		if err != nil {
			c.Error(&errors.InvalidInputError{
				Message: "Invalid chipID parameter",
				Details: map[string]any{
					"original": err.Error(),
					"chip_id":  chipIDString,
				},
			})
			return
		}
		chipID := int32(val)
		fromTime, fromParsingErr := time.Parse(TIME_LAYOUT, from)
		toTime, toParsingErr := time.Parse(TIME_LAYOUT, to)
		if (fromParsingErr != nil) || (to != "" && toParsingErr != nil) {
			c.Error(&errors.InvalidInputError{
				Message: "Missing or invalid query parameter",
				Details: map[string]any{
					"from":            from,
					"to":              to,
					"expected_layout": "YYYY-MM-DD",
				},
			})
			return
		}
		tels, srvErr := service.GetRange(chipID, fromTime, toTime)
		if srvErr != nil {
			c.Error(srvErr)
			return
		}

		c.JSON(http.StatusOK, http_mappers.ToTelemetryRangeResponse(tels))
	}
}
