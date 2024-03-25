package services

import (
	"fmt"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/ports"
)

type IngestorService struct {
	Telemetry ports.TelemetryPort
	Repo      ports.RepoPort
}

func (s *IngestorService) HandleTelemetry(data []byte) entities.Telemetry {
	telemetry, extractError := s.Telemetry.ExtractData(data)
	if extractError != nil {
		fmt.Println("error extracting telemetry")
		return entities.Telemetry{}
	}
	repoError := s.Repo.Save(telemetry.Id, telemetry)
	if repoError != nil {
		fmt.Println("error saving telemetry")
	}
	return telemetry
}
