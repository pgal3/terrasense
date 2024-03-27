package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/PaoloEG/terrasense/internal/core/services"
	"github.com/PaoloEG/terrasense/internal/infrastructure/adapters/influxdb"
	protobuf_adapter "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/protobuf"
	mqtt_handler "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/mqtt"
	"github.com/joho/godotenv"
)

const MQTT_TOPIC = "terrasense/+/measurements"

func main() {
	isProd := flag.Bool("prod", false, "Set environment as PROD")
	fmt.Println("Program Starting...")
	flag.Parse()
	if(!*isProd){
		fmt.Println("Loading .env")
		currentPath, _ := os.Getwd()
		godotenv.Load(currentPath+"/.env")
	}
	ingestorService := &services.IngestorService{
		Telemetry: protobuf_adapter.New(),
		Repo:      influxdb.New(influxdb.WithEnvConfig()),
	}
	mqtt := mqtt_handler.New(mqtt_handler.WithEnvConfig())
	mqtt.Start(MQTT_TOPIC, ingestorService.CreateTelemetryHandler())

	// Wait for a signal to exit the program gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	mqtt.Disconnect()
}
