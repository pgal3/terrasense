package influxdb

import (
	"context"
	"fmt"
	"log"

	// "time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type Client struct {
	client influxdb2.Client
	config influxDBConfig
}

type Option func(*Client)

func New(options ...Option) *Client {
	client := &Client{}
	for _, opt := range options {
		opt(client)
	}
	return client
}

func (c *Client) Save(id string, telemetry entities.Telemetry) error {
	tags := map[string]string{
		"version":       "1",
		"sensorID":      fmt.Sprint(telemetry.ChipID),
		"measurementID": telemetry.Id,
	}
	fields := map[string]any{
		"temperature":  telemetry.Measurements.Temperature,
		"soilMoisture": telemetry.Measurements.SoilMoisture,
		"humidity":     telemetry.Measurements.Humidity,
		"pressure":     telemetry.Measurements.Pressure,
		"altitude":     telemetry.Measurements.Altitude,
	}
	point := write.NewPoint("terrasense_measurements", tags, fields, telemetry.Timestamp)
	writeAPI := c.client.WriteAPIBlocking(c.config.Org, c.config.Bucket)
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		fmt.Println("error in saving point to influxDB")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (c *Client) GetLatest(chipID string)(entities.Telemetry, error){
	queryAPI := c.client.QueryAPI(c.config.Org)
	query := fmt.Sprintf(`from(bucket: %s)
    |> range(start: 0)
    |> filter(fn: (r) => r["_measurement"] == "terrasense_measurements")
	|> filter(fn: (r) => r["sensorID"] == %s)
    |> last()`, c.config.Bucket, chipID)
	result, err := queryAPI.Query(context.Background(), query)
	log.Println(result)
	log.Println(err)
	return entities.Telemetry{}, nil
}

// func (c *Client) GetRange(from time.Time, to time.Time)([]entities.Telemetry, error){
	
// }
