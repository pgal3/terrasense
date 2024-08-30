package pb_adapter

import (
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func CreateMeasurementsDecoder(inChannel <- chan []byte, outChannel chan <- entities.Telemetry){
	go func(){
		for data := range inChannel {
			decodedData, _ := getMeasurements(data)
			outChannel <- createTelemetry(decodedData)
		}
	}()
}

func getMeasurements(data []byte)(*Measurements, error) {
	pbPayload := &Measurements{}
	if err := proto.Unmarshal(data, pbPayload); err != nil {
		return pbPayload, err
	}
	return pbPayload, nil
}

func createTelemetry(data *Measurements) entities.Telemetry {
	return entities.Telemetry{
		ID:        uuid.NewString(),
		Version:   "1", //TODO: IMPLEMENT VERSION IN THE PROTO
		ChipID:    data.ChipID,
		Timestamp: time.Now().UTC(),
		Measurement: vo.NewMeasurement(
			data.Temperature,
			data.Soil,
			data.Humidity,
			data.Pressure,
			data.Altitude,
		),
	}
}
