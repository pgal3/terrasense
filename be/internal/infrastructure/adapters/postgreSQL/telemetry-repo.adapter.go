package pg_adapter

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	pg_mappers "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/mappers"
	pg_models "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Client struct {
	pgdb 	*sql.DB
	bundb 	*bun.DB
	config  PostgreSQLConfig
	ctx		context.Context
}

func New(config PostgreSQLConfig, ctx context.Context) *Client {
	db := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%s",config.Url,config.Port)),
		pgdriver.WithUser(config.User),
		pgdriver.WithPassword(config.Pwd),
		pgdriver.WithInsecure(true),
		))
	bundb := bun.NewDB(db, pgdialect.New())
	bundb.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(false),
		// bundebug.FromEnv("BUNDEBUG"),
		bundebug.WithEnabled(true),
	))

	pg_models.SetUpMeasurementTable(bundb)

	return &Client{
		pgdb: db,
		bundb: bundb,
		config: config,
		ctx: ctx,
	}
}

func(c *Client) Save(id string, telemetry entities.Telemetry) error {
	dbMeasurement := pg_mappers.MapMeasurement(telemetry)
	_, err := c.bundb.NewInsert().Model(&dbMeasurement).Exec(c.ctx)
	if err != nil {
		return &errors.InternalServerError{
			Message: "Error in saving telemetry in DB",
			OriginalError: err.Error(),
		}
	}
	return nil
}

func(c *Client) GetLatest(chipID int32)(entities.Telemetry, error){
	measurement := pg_models.Measurement{}
	err := c.bundb.NewSelect().Model(&measurement).Where("chip_id = ?", chipID).Limit(1).OrderExpr("timestamp DESC").Scan(c.ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return entities.Telemetry{}, &errors.NotFoundError{
				Message: "No measurements found",
			}
		}
		return entities.Telemetry{}, &errors.InternalServerError{
			Message: "Error in running DB query",
			OriginalError: err.Error(),
		}
	}
	telemetry := pg_mappers.MapTelemetry(measurement)
	return telemetry, nil
}

func(c *Client) Close(){
	c.bundb.Close()
	c.pgdb.Close()
}