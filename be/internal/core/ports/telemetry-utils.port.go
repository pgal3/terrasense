package ports

import "github.com/PaoloEG/terrasense/internal/core/domain/entities"

type TelemetryUtilsPort interface {
	ExtractData([]byte) (entities.Telemetry, error)
}
