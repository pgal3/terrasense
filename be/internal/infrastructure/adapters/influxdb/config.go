package influxdb

import "os"

type influxDBConfig struct {
	Url string
	Token string
	Org string
	Bucket string
	BatchSize uint

	MeasurementsName string
}

var WithEnvConfig = influxDBConfig{
	Url: os.Getenv("INFLUXDB_URL"),
	Token: os.Getenv("INFLUXDB_TOKEN"),
	Org: os.Getenv("INFLUXDB_ORG"),
	Bucket: os.Getenv("INFLUXDB_BUCKET"),
	BatchSize: 1,
	MeasurementsName: "terrasense_measurements",
}