package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/PaoloEG/terrasense/internal/core/services"
	"github.com/PaoloEG/terrasense/internal/infrastructure/adapters/influxdb"
	protobuf_adapter "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/protobuf"
	mqtt_lib "github.com/PaoloEG/terrasense/lib/mqtt"
	"github.com/joho/godotenv"
)

func main() {
	keepAlive := make(chan os.Signal)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGTERM)
	godotenv.Load("../.env")
	fmt.Println("Program Starting...")
	config := &mqtt_lib.MqttOptions{
		Host:     os.Getenv("MQTT_HOST"),
		Port:     os.Getenv("MQTT_PORT"),
		ClientId: os.Getenv("MQTT_CLIENTID"),
		Username: os.Getenv("MQTT_USERNAME"),
		Password: os.Getenv("MQTT_PWD"),
	}
	mqtt := mqtt_lib.New(*config)
	ingestor := &services.IngestorService{
		Telemetry: protobuf_adapter.New(),
		Repo:      influxdb.New(),
	}
	funz := func(input []byte) {
		ingestor.HandleTelemetry(input)
	}
	mqtt.Sub("/terrasense/+/measurements", funz)
	<-keepAlive
}
