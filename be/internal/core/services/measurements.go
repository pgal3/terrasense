package services

import (
	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/ports"
)

type MeasurementsService struct {
	Repo      ports.TelemetryRepoPort
}

func (s * MeasurementsService) GetLatestMeasurement(chipID string)(entities.Telemetry, error) {
	return s.Repo.GetLatest(chipID)
}