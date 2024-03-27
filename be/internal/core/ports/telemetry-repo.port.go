package ports

import (
	// "time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
)

type TelemetryRepoPort interface {
	Save(string, entities.Telemetry) error
	GetLatest(chipID string)(entities.Telemetry, error)
	// GetRange(from time.Time, to time.Time)([]entities.Telemetry, error)
	// Delete(id string) error
}
