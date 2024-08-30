package services_mappers

import "github.com/PaoloEG/terrasense/internal/core/domain/entities"

func FromPairingsToMap(pairs []entities.Pairing) map[string]map[string]any {
	keyValue := make(map[string]map[string]any)
	for _, val := range pairs {
		keyValue[fmt.Sprintf("%s-%d",val.UserID, val.ChipID)] = map[string]any{
			"threshold": val.Settings.LowLevelThreshold(),
			"notifyMe": val.Settings.NotifyMe(),
		}
	}
	return keyValue
}