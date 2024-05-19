package services

import (
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/ports"
)

type MeasurementsService struct {
	repo ports.TelemetryRepoPort
}

func NewMeasurementsService(telemetryRepo ports.TelemetryRepoPort) *MeasurementsService {
	return &MeasurementsService{
		repo: telemetryRepo,
	}
}

func (s *MeasurementsService) GetLatestMeasurement(chipID int32) (entities.Telemetry, error) {
	return s.repo.GetLatest(chipID)
}

func (s *MeasurementsService) GetRange(chipID int32, from time.Time, to time.Time) ([]entities.Telemetry, error) {
	return s.repo.GetRange(chipID, from, to)
}
