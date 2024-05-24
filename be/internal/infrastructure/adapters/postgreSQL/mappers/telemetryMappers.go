package pg_mappers

import (
	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
	pg_models "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/models"
)

func ToMeasurementModel(telemetry entities.Telemetry) pg_models.Measurement {
	return pg_models.Measurement{
		ID:           telemetry.ID,
		ChipID:       telemetry.ChipID,
		Timestamp:    telemetry.Timestamp,
		Version:      telemetry.Version,
		Temperature:  telemetry.Measurement.Temperature(),
		Humidity:     telemetry.Measurement.Humidity(),
		Altitude:     telemetry.Measurement.Altitude(),
		SoilMoisture: telemetry.Measurement.SoilMoisture(),
		Pressure:     telemetry.Measurement.Pressure(),
	}
}

func ToTelemetryEntity(ms pg_models.Measurement) entities.Telemetry {
	return entities.Telemetry{
		ID:        ms.ID,
		ChipID:    ms.ChipID,
		Version:   ms.Version,
		Timestamp: ms.Timestamp,
		Measurement: vo.NewMeasurement(
			ms.Temperature,
			ms.SoilMoisture,
			ms.Humidity,
			ms.Pressure,
			ms.Altitude,
		),
	}
}

func ToTelemetryEntities(ms []pg_models.Measurement) []entities.Telemetry {
	telemetries := []entities.Telemetry{}
	for _, val := range ms {
		telemetries = append(telemetries, entities.Telemetry{
			ID:        val.ID,
			ChipID:    val.ChipID,
			Version:   val.Version,
			Timestamp: val.Timestamp,
			Measurement: vo.NewMeasurement(
				val.Temperature,
				val.SoilMoisture,
				val.Humidity,
				val.Pressure,
				val.Altitude,
			),
		})
	}
	return telemetries
}
