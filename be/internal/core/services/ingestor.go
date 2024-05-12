package services

import (
	"log"

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

func (s *IngestorService) TelemetryHandler(data []byte){
		telemetry, extractError := s.telemetry.GetTelemetry(data)
		if extractError != nil {
			log.Printf("error extracting telemetry: %s", extractError.Error())
		}
		repoError := s.repo.Save(telemetry.Id, telemetry)
		if repoError != nil {
			log.Printf("error saving telemetry: %s", repoError.Error())
		}
}
