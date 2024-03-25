package ports

import "github.com/PaoloEG/terrasense/internal/core/domain/entities"

type RepoPort interface {
	Save(string, entities.Telemetry) error
	// Get(id string)(entities.Telemetry, error)
}
