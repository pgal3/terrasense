package entities

import vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"

type Pairing struct {
	UserID		string
	ChipID		int32
	Settings    vo.PairingSettings
}