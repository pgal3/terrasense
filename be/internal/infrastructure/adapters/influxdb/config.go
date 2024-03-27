package influxdb

type influxDBConfig struct {
	Url string
	Token string
	Org string
	Bucket string
	BatchSize uint
}