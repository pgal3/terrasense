package services

import (
	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/ports"
)

type MeasurementsService struct {
	repo      ports.TelemetryRepoPort
}

func NewMeasurementsService(telemetryRepo ports.TelemetryRepoPort) *MeasurementsService {
	return &MeasurementsService{
		repo: telemetryRepo,
	}
}

func (s * MeasurementsService) GetLatestMeasurement(chipID int32)(entities.Telemetry, error) {
	return s.repo.GetLatest(chipID)
}