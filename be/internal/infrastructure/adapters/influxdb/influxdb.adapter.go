package influxdb

import (
	"context"
	"fmt"
	"os"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type Client struct {
	client influxdb2.Client
}

func New() *Client {
	return &Client{
		client: influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN")),
	}
}

func (c *Client) Save(id string, telemetry entities.Telemetry) error {
	writeApi := c.client.WriteAPIBlocking(os.Getenv("INFLUXDB_ORG"), os.Getenv("INFLUXDB_BUCKET"))
	tags := map[string]string{
		"version":       "1",
		"sensorId":      fmt.Sprint(telemetry.ChipID),
		"measurementId": telemetry.Id,
	}
	fields := map[string]any{
		"temperature":  telemetry.Measurements.Temperature,
		"soilMoisture": telemetry.Measurements.SoilMoisture,
		"humidity":     telemetry.Measurements.Humidity,
		"pressure":     telemetry.Measurements.Pressure,
		"altitude":     telemetry.Measurements.Altitude,
	}
	point := write.NewPoint("terrasense_measurements", tags, fields, telemetry.Timestamp)
	if err := writeApi.WritePoint(context.Background(), point); err != nil {
		fmt.Println("error in saving point to influxDB")
		fmt.Println(err.Error())
		return err
	}
	return nil
}
