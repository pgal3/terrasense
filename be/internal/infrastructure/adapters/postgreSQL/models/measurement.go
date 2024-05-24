package pg_models

import "time"

type Measurement struct {
	ID           string    `gorm:"primaryKey;not null"`
	Timestamp    time.Time `gorm:"index:meas_idx;not null;default:current_timestamp"`
	ChipID       int32     `gorm:"index:meas_idx;not null"`
	Version      string    `gorm:"version;not null;default:'1'"`
	Temperature  float32
	Humidity     float32
	Altitude     float32
	Pressure     float32
	SoilMoisture float32
}
