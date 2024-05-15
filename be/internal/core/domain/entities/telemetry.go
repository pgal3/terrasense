package entities

import (
	"time"

	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
)

type Telemetry struct {
	ID          string
	ChipID      int32
	Version     string
	Timestamp   time.Time
	Measurement vo.Measurement
}
