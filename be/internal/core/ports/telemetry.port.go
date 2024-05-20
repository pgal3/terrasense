package ports

import "github.com/PaoloEG/terrasense/internal/core/domain/entities"

type TelemetryPort interface {
	GetTelemetry([]byte) (entities.Telemetry, error)
}
