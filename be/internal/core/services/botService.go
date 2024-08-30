package services

import (
	"fmt"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	"github.com/PaoloEG/terrasense/internal/core/ports"
	services_mappers "github.com/PaoloEG/terrasense/internal/core/services/mappers"
)

type BotService struct {
	pairRepo ports.PairingRepoPort
	telemetryRepo ports.TelemetryRepoPort
}

func NewBotService(pairingRepoPort ports.PairingRepoPort, telemetryRepoPort ports.TelemetryRepoPort) *BotService {
	return &BotService{
		pairRepo: pairingRepoPort,
		telemetryRepo: telemetryRepoPort,
	}
}

func(s *BotService) PairDevice(pairing entities.Pairing) error {
	return s.pairRepo.PairDevice(pairing)
}

func(s *BotService) DeletePairing(userID string, chipID int32) error {
	return s.pairRepo.DeletePair(userID, chipID)
}

func(s *BotService) GetPairedDevice(userID string)([]entities.Pairing, error){
	return s.pairRepo.GetPairings(userID)
}

func(s *BotService) UpdateSettings(pairing entities.Pairing) error {
	return s.pairRepo.UpdatePairing(pairing.UserID, pairing.ChipID, pairing.Settings)
}

func(s *BotService) GetLatestMeasurements(userID string, chipID int32) (entities.Telemetry, error) {
	//TODO: add a cache layer
	userPairings, err := s.pairRepo.GetPairings(userID)
	if err != nil {
		return entities.Telemetry{}, err
	}
	pairs := services_mappers.FromPairingsToMap(userPairings)
	if pairs[fmt.Sprintf("%s-%d",userID, chipID)] == nil {
		return entities.Telemetry{}, &errors.ForbiddenError{
			Message: fmt.Sprintf("UserID %s and ChipID %d are not paired", userID, chipID),
		}
	}
	return s.telemetryRepo.GetLatest(chipID)
}