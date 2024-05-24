package http_mappers

import (
	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	http_res "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/http/res"
	http_dtos "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/http/res/dto"
)

func ToTelemetryResponse(telemetry entities.Telemetry) http_res.LatestTelemetryResponse {
	return http_res.LatestTelemetryResponse{
		Timestamp: telemetry.Timestamp.UTC(),
		ID:        telemetry.ID,
		Measurement: http_dtos.MeasurementDto{
			Temperature:  telemetry.Measurement.Temperature(),
			SoilMoisture: telemetry.Measurement.SoilMoisture(),
			Humidity:     telemetry.Measurement.Humidity(),
			Pressure:     telemetry.Measurement.Pressure(),
			Altitude:     telemetry.Measurement.Altitude(),
		},
	}
}

func ToTelemetryRangeResponse(telemetries []entities.Telemetry) http_res.RangeTelemetryResponse {
	res := http_res.RangeTelemetryResponse{}
	for _, telemetry := range telemetries {
		res = append(res, http_res.LatestTelemetryResponse{
			Timestamp: telemetry.Timestamp.UTC(),
			ID:        telemetry.ID,
			Measurement: http_dtos.MeasurementDto{
				Temperature:  telemetry.Measurement.Temperature(),
				SoilMoisture: telemetry.Measurement.SoilMoisture(),
				Humidity:     telemetry.Measurement.Humidity(),
				Pressure:     telemetry.Measurement.Pressure(),
				Altitude:     telemetry.Measurement.Altitude(),
			},
		})
	}
	return res
}
