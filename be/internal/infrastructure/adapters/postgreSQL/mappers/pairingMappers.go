package pg_mappers

import (
	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
	pg_models "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/models"
)

func ToPairingModel(pairing entities.Pairing) pg_models.Pairing {
	return pg_models.Pairing{
		UserID: pairing.UserID,
		ChipID: pairing.ChipID,
		NotifyMe: pairing.Settings.NotifyMe(),
		LowLevelThreshold: pairing.Settings.LowLevelThreshold(),
	}
}

func ToPairingEntities(pairings []pg_models.Pairing) []entities.Pairing {
	pairs := []entities.Pairing{}
	for _, val := range pairings {
		pairs = append(pairs, entities.Pairing{
			UserID: val.UserID,
			ChipID: val.ChipID,
			Settings: vo.NewPairSettings(
				val.NotifyMe,
				val.LowLevelThreshold,
			),
		})
	}
	return pairs
}