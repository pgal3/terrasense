package http_dtos

type MeasurementDto struct {
	Temperature  float32 `json:"temperature,omitempty"`
	SoilMoisture float32 `json:"soilMoisture,omitempty"`
	Humidity     float32 `json:"humidity,omitempty"`
	Pressure     float32 `json:"pressure,omitempty"`
	Altitude     float32 `json:"altitude,omitempty"`
}