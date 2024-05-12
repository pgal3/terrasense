package influxdb

import (
	"context"
	"fmt"
	"log"

	// "time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	mappers "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/influxdb/mappers"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type Client struct {
	client influxdb2.Client
	config influxDBConfig
}

func New(config influxDBConfig) *Client {
	client := &Client{
		client: influxdb2.NewClientWithOptions(config.Url, config.Token, influxdb2.DefaultOptions().SetUseGZip(true).SetBatchSize(config.BatchSize)),
		config: config,
	}
	return client
}

func (c *Client) Save(id string, telemetry entities.Telemetry) error {
	tags := mappers.ToTelemetryTags(telemetry)
	fields := mappers.ToTelemetryFields(telemetry.Measurements)
	point := write.NewPoint(c.config.MeasurementsName, tags, fields, telemetry.Timestamp)
	writeAPI := c.client.WriteAPIBlocking(c.config.Org, c.config.Bucket)
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Println("Error in saving point to influxDB: ", err.Error())
		return err
	}
	return nil
}

func (c *Client) GetLatest(chipID string)(entities.Telemetry, error){
	queryAPI := c.client.QueryAPI(c.config.Org)
	query := fmt.Sprintf(`from(bucket: %s)
    |> range(start: 0)
    |> filter(fn: (r) => r["sensorID"] == %s and r["_measurement"] == "terrasense_measurements")
	|> group(columns: ["_field"])
    |> last()`, c.config.Bucket, chipID) //TODO: sanitize this values first
	result, err := queryAPI.Query(context.Background(), query)
	log.Println(result)
	log.Println(err)
	return entities.Telemetry{}, nil
}

// func (c *Client) GetRange(from time.Time, to time.Time)([]entities.Telemetry, error){
	/*
	//Possible query:
	`from(bucket: "terrasense_bucket")
  	|> range(start: -1h, stop: -1m)
	|> filter(fn: (r) => r.sensorID == "2703511" and r._measurement == "terrasense_measurements")
  	|> group(columns: ["measurementID"])`
	*/
// }
