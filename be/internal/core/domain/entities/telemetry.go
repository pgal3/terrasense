package entities

import (
	"time"

	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
)

type Telemetry struct {
	Id           string
	ChipID       int32
	Name         string
	Timestamp    time.Time
	Measurements vo.Measurement
}
