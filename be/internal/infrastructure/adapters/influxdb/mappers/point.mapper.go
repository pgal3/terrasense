package influxdb_mappers

import (
	"fmt"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
)

func ToTelemetryTags(telemetry entities.Telemetry) map[string]string {
	return map[string]string{
		"version":       telemetry.Version,
		"sensorID":      fmt.Sprint(telemetry.ChipID),
		"measurementID": telemetry.Id,
	}
}

func ToTelemetryFields(measurements vo.Measurement) map[string]any{
	return map[string]any{
		"temperature":  measurements.Temperature(),
		"soilMoisture": measurements.SoilMoisture(),
		"humidity":     measurements.Humidity(),
		"pressure":     measurements.Pressure(),
		"altitude":     measurements.Altitude(),
	}
}