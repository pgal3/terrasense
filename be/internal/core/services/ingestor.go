package services

import (
	"log"

	"github.com/PaoloEG/terrasense/internal/core/ports"
)

type IngestorService struct {
	Telemetry ports.TelemetryUtilsPort
	Repo      ports.TelemetryRepoPort
}

func (s *IngestorService) CreateTelemetryHandler() func([]byte){
	return func(data []byte){
		telemetry, extractError := s.Telemetry.ExtractData(data)
		if extractError != nil {
			log.Printf("error extracting telemetry: %s", extractError.Error())
		}
		repoError := s.Repo.Save(telemetry.Id, telemetry)
		if repoError != nil {
			log.Printf("error saving telemetry: %s", repoError.Error())
		}
	}
}
