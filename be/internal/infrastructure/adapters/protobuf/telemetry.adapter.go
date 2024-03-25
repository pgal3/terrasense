package protobuf_adapter

import (
	// "encoding/base64"
	"fmt"
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type Adapter struct{}

func New() *Adapter {
	return &Adapter{}
}

func (a *Adapter) ExtractData(data []byte) (entities.Telemetry, error) {
	pbPayload := &Measurements{}
	// decodedData, _ := base64.StdEncoding.DecodeString(string(data[:]))
	if err := proto.Unmarshal(data, pbPayload); err != nil {
		fmt.Println("Error unmarshalling the payload:")
		fmt.Println(err.Error())
		return entities.Telemetry{}, err
	}

	return entities.Telemetry{
		Id:        uuid.NewString(),
		ChipID:    pbPayload.ChipID,
		Name:      "",
		Timestamp: time.Now(),
		Measurements: vo.Measurement{
			Temperature:  pbPayload.Temperature,
			SoilMoisture: pbPayload.Soil,
			Humidity:     pbPayload.Humidity,
			Pressure:     pbPayload.Pressure,
			Altitude:     pbPayload.Altitude,
		},
	}, nil
}
