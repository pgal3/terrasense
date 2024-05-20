package pg_adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	pg_mappers "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/mappers"
	pg_models "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db     *gorm.DB
	config PostgreSQLConfig
	ctx    context.Context
}

func New(config PostgreSQLConfig, ctx context.Context) *Client {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Url, config.User, config.Pwd, config.DBName, config.Port)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Setup Tables
	db.AutoMigrate(&pg_models.Measurement{})

	return &Client{
		db:     db,
		config: config,
		ctx:    ctx,
	}
}

func (c *Client) Save(id string, telemetry entities.Telemetry) error {
	dbMeasurement := pg_mappers.MapMeasurement(telemetry)
	res := c.db.Create(&dbMeasurement)
	if res.Error != nil {
		return &errors.InternalServerError{
			Message:       "Error in saving telemetry in DB",
			OriginalError: res.Error.Error(),
		}
	}
	return nil
}

func (c *Client) GetLatest(chipID int32) (entities.Telemetry, error) {
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
	telemetry := pg_mappers.MapTelemetry(measurement)
	return telemetry, nil
}

func (c *Client) GetRange(chipID int32, from time.Time, to time.Time) ([]entities.Telemetry, error) {
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
	telemetries := pg_mappers.MapTelemetries(measurements)
	return telemetries, nil
}

func (c *Client) Close() {
	db, _ := c.db.DB()
	db.Close()
}
