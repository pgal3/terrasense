package pg_models

import (
	"context"
	"log"
	"time"

	"github.com/uptrace/bun"
)

type Measurement struct {
	bun.BaseModel `bun:"table:measurements"`

	ID           string    `bun:"id,pk,type:uuid,notnull"`
	Timestamp    time.Time `bun:"timestamp,nullzero,notnull,default:current_timestamp"`
	ChipID       int32     `bun:"chip_id"`
	Version 	 string    `bun:"version,notnull,default:'1'"`
	Temperature  float32   `bun:"temperature"`
	Humidity     float32   `bun:"humidity"`
	Altitude     float32   `bun:"altitude"`
	Pressure     float32   `bun:"pressure"`
	SoilMoisture float32   `bun:"soil_moisture"`
}

func SetUpMeasurementTable(db *bun.DB){
	ctx := context.Background()
	_, tableErr := db.NewCreateTable().Model((*Measurement)(nil)).IfNotExists().Exec(ctx)
	if tableErr != nil {
		log.Panic(tableErr)
	}
	_, indexErr := db.NewCreateIndex().IfNotExists().Model((*Measurement)(nil)).Index("meas_idx").Column("timestamp","chip_id").Exec(ctx)
	if indexErr != nil {
		log.Panic(indexErr)
	}
}
