package mqtt_handler

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttHandler struct {
	client mqtt.Client
}

type Option func(*MqttHandler)

func New(options ...Option) *MqttHandler {
	client := &MqttHandler{}
	for _, opt := range options {
		opt(client)
	}
	return client
}

func (mq *MqttHandler) Start(subTopic string, onMessageReceive func([]byte)) {
	if token := mq.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error in connecting MQTT client: %s", token.Error().Error())
	}
	subToken := mq.client.Subscribe(subTopic, 0, func(cl mqtt.Client, m mqtt.Message) {
		log.Printf("Message received from topic %s", m.Topic())
		onMessageReceive(m.Payload())
	})
	subToken.Wait()
	if subToken.Error() != nil {
		mq.client.Disconnect(100)
		log.Fatalf("Failed to subscribe to topic: %s", subToken.Error().Error())
	}
	log.Printf("Subscribed to topic %s", subTopic)
}

func (mq *MqttHandler) Disconnect() {
		log.Println("Disconnecting MQTT client")
		mq.client.Disconnect(300)
}
