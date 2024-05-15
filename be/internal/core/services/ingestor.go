package services

import (
	"github.com/PaoloEG/terrasense/internal/core/ports"
)

type IngestorService struct {
	telemetry ports.TelemetryPort
	repo      ports.TelemetryRepoPort
}

func NewIngestorService(telemetryPort ports.TelemetryPort, telemetryRepo ports.TelemetryRepoPort) *IngestorService {
	return &IngestorService{
		telemetry: telemetryPort,
		repo: telemetryRepo,
	}
}

func (s *IngestorService) TelemetryHandler(data []byte) error {
		telemetry, extractError := s.telemetry.GetTelemetry(data)
		if extractError != nil {
			return extractError
		}
		repoError := s.repo.Save(telemetry.ID, telemetry)
		if repoError != nil {
			return repoError
		}
		return nil
}
