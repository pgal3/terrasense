package pb_adapter

import (
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type TelemetryPortAdapter struct{}

func New() *TelemetryPortAdapter {
	return &TelemetryPortAdapter{}
}

func (t *TelemetryPortAdapter) GetTelemetry(data []byte) (entities.Telemetry, error) {
	pbPayload := &Measurements{}
	if err := proto.Unmarshal(data, pbPayload); err != nil {
		return entities.Telemetry{}, &errors.InternalServerError{
			Message:       "Error converting incoming protobuf",
			OriginalError: err.Error(),
		}
	}

	return entities.Telemetry{
		ID:        uuid.NewString(),
		Version:   "1", //TODO: IMPLEMENT VERSION IN THE PROTO
		ChipID:    pbPayload.ChipID,
		Timestamp: time.Now().UTC(),
		Measurement: vo.NewMeasurement(
			pbPayload.Temperature,
			pbPayload.Soil,
			pbPayload.Humidity,
			pbPayload.Pressure,
			pbPayload.Altitude,
		),
	}, nil
}
