package pg_adapter

import (
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	pg_mappers "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/mappers"
	pg_models "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/models"
	"gorm.io/gorm"
)

type TelemetryRepoAdapter struct {
	db *gorm.DB
}

func NewTelemetryRepoAdapter(dbClient *gorm.DB) *TelemetryRepoAdapter{
	return &TelemetryRepoAdapter{
		db: dbClient,
	}
}

func (c *TelemetryRepoAdapter) Save(id string, telemetry entities.Telemetry) error {
	dbMeasurement := pg_mappers.ToMeasurementModel(telemetry)
	res := c.db.Create(&dbMeasurement)
	if res.Error != nil {
		return &errors.InternalServerError{
			Message:       "Error in saving telemetry in DB",
			OriginalError: res.Error.Error(),
		}
	}
	return nil
}

func (c *TelemetryRepoAdapter) GetLatest(chipID int32) (entities.Telemetry, error) {
	measurement := pg_models.Measurement{}
	res := c.db.Order("timestamp DESC").Where("chip_id = ?", chipID).Last(&measurement)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return entities.Telemetry{}, &errors.NotFoundError{
				Message: "No measurements found",
			}
		}
		return entities.Telemetry{}, &errors.InternalServerError{
			Message:       "Error in running DB query",
			OriginalError: res.Error.Error(),
		}
	}
	telemetry := pg_mappers.ToTelemetryEntity(measurement)
	return telemetry, nil
}

func (c *TelemetryRepoAdapter) GetRange(chipID int32, from time.Time, to time.Time) ([]entities.Telemetry, error) {
	measurements := []pg_models.Measurement{}
	query := c.db.Where("chip_id = ?", chipID)
	if !to.IsZero() {
		query.Where("timestamp <= ?", to)
	}
	res := query.Order("timestamp DESC").Find(&measurements)
	if res.Error != nil {
		return []entities.Telemetry{}, &errors.InternalServerError{
			Message:       "Error in running DB query",
			OriginalError: res.Error.Error(),
		}
	}
	telemetries := pg_mappers.ToTelemetryEntities(measurements)
	return telemetries, nil
}

func (c *TelemetryRepoAdapter) Delete(id string) error {
	delete := c.db.Delete(&pg_models.Measurement{}, id)
	if delete.Error != nil {
		return &errors.InternalServerError{
			Message:       "Error in running DB delete",
			OriginalError: delete.Error.Error(),
		}
	}
	return nil
}


