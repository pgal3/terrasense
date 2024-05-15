package http_res

import (
	"time"

	http_dtos "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/http/res/dto"
)

type LatestTelemetryResponse struct {
	ID          string                   `json:"id,omitempty"`
	Timestamp   time.Time                `json:"timestamp,omitempty"`
	Measurement http_dtos.MeasurementDto `json:"measurement,omitempty"`
}
