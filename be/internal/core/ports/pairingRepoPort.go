package ports

import (
	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
)

type PairingRepoPort interface {
	PairDevice(pairing entities.Pairing) error
	GetPairings(userID string)([]entities.Pairing, error)
	DeletePair(userID string, chipID int32) error
	UpdatePairing(userID string, chipID int32, settings vo.PairingSettings) error
}