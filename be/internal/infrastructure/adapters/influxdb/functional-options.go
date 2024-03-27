package influxdb

import (
	"os"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func WithEnvConfig() func(*Client) {
	return func(c *Client){
		c.config.Url = os.Getenv("INFLUXDB_URL")
		c.config.Token = os.Getenv("INFLUXDB_TOKEN")
		c.config.Org = os.Getenv("INFLUXDB_ORG")
		c.config.Bucket = os.Getenv("INFLUXDB_BUCKET")
		c.config.BatchSize = 1
		c.client = influxdb2.NewClientWithOptions(c.config.Url, c.config.Token, influxdb2.DefaultOptions().SetUseGZip(true).SetBatchSize(c.config.BatchSize))
	}
}