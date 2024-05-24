package mqtt_hdl

import (
	"log"

	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	"github.com/PaoloEG/terrasense/internal/core/services"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttHandler struct {
	client  mqtt.Client
	service *services.IngestorService
}

type Option func(*MqttHandler)

func New(options ...Option) *MqttHandler {
	client := &MqttHandler{}
	for _, opt := range options {
		opt(client)
	}
	return client
}

func (mq *MqttHandler) Start() error {
	if token := mq.client.Connect(); token.Wait() && token.Error() != nil {
		return &errors.UnknownError{
			Message: "Error in connecting MQTT client",
			Details: map[string]any{
				"original": token.Error().Error(),
			},
		}
	}
	opts := mqtt.ClientOptionsReader(mq.client.OptionsReader())
	log.Println("Connected to MQTT - ClientID:",opts.ClientID())
	return nil
}

func (mq *MqttHandler) Subscribe(subTopic string) error {
	subToken := mq.client.Subscribe(subTopic, 0, func(cl mqtt.Client, m mqtt.Message) {
		log.Println("Message received from topic", m.Topic())
		go mq.service.TelemetryHandler(m.Payload())
	})
	subToken.Wait()
	if subToken.Error() != nil {
		mq.client.Disconnect(100)
		log.Fatalf("Failed to subscribe to topic: %s", subToken.Error().Error())
		return subToken.Error()
	}
	log.Printf("Subscribed to topic %s", subTopic)
	return nil
}

func (mq *MqttHandler) Disconnect() {
	log.Println("Disconnecting MQTT client")
	mq.client.Disconnect(300)
}
