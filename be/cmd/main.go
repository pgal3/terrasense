package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/PaoloEG/terrasense/internal/core/services"
	pg_adapter "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL"
	protobuf_adapter "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/protobuf"
	http_handler "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/http"
	mqtt_handler "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/mqtt"
	"github.com/joho/godotenv"
)

const MQTT_TOPIC = "terrasense/+/measurements"

func main() {
	// ======= READ FLAG =======
	isProd := flag.Bool("prod", false, "Set environment as PROD")
	fmt.Println("Program Starting...")
	flag.Parse()
	if !*isProd {
		fmt.Println("Loading .env")
		currentPath, _ := os.Getwd()
		godotenv.Load(currentPath + "/local.env")
	}

	// ======= DI =======
	dbClient := pg_adapter.NewDBClient(pg_adapter.WithEnvConfig)
	telemetryRepo := pg_adapter.NewTelemetryRepoAdapter(dbClient)

	ingestorService := services.NewIngestorService(
		protobuf_adapter.New(),
		telemetryRepo,
	)
	measurementService := services.NewMeasurementsService(telemetryRepo)

	// ======= Initialize Handlers =======
	mqtt := mqtt_handler.New(mqtt_handler.WithEnvConfig(ingestorService))
	http := http_handler.New(measurementService, *isProd)

	// ======= START MQTT HANDLER =======
	mqtt.Start()               //TODO: manage errors
	mqtt.Subscribe(MQTT_TOPIC) //TODO: manage errors

	// ======= START HTTP HANDLER =======
	http.Start("3000")

	// Wait for a signal to exit the program gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	mqtt.Disconnect()
}
