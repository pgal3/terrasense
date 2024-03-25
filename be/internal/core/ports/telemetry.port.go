package ports

import "github.com/PaoloEG/terrasense/internal/core/domain/entities"

type TelemetryPort interface {
	ExtractData([]byte) (entities.Telemetry, error)
}
