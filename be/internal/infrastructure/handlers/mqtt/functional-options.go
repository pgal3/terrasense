package mqtt_hdl

import (
	"fmt"
	"os"

	"github.com/PaoloEG/terrasense/internal/core/services"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func WithEnvConfig(srv *services.IngestorService) func(*MqttHandler) {
	return func(h *MqttHandler) {
		config := mqtt.NewClientOptions()
		config.AddBroker(fmt.Sprintf("ssl://%s:%s/mqtt", os.Getenv("MQTT_HOST"), os.Getenv("MQTT_PORT")))
		config.SetClientID(os.Getenv("MQTT_CLIENTID"))
		config.SetUsername(os.Getenv("MQTT_USERNAME"))
		config.SetPassword(os.Getenv("MQTT_PWD"))
		config.SetAutoReconnect(true)
		h.client = mqtt.NewClient(config)
		h.service = srv
	}
}
