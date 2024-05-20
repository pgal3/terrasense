package ports

import (
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
)

type TelemetryRepoPort interface {
	Save(id string, telemetry entities.Telemetry) error
	GetLatest(chipID int32) (entities.Telemetry, error)
	GetRange(chipID int32, from time.Time, to time.Time) ([]entities.Telemetry, error)
	Delete(id string) error
}
