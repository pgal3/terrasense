package mqtt_hdl

import (
	"fmt"
	"math/rand"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttBuilder struct {
	config          *mqtt.ClientOptions
	rcvChannel      chan <- []byte
}

func NewBuilder() MqttBuilder {
	return MqttBuilder{
		config:     mqtt.NewClientOptions(),
		rcvChannel: nil,
	}
}

func (b *MqttBuilder) SetEnvConfig() {
	b.config.AddBroker(fmt.Sprintf("ssl://%s:%s/mqtt", os.Getenv("MQTT_HOST"), os.Getenv("MQTT_PORT")))
	b.config.SetClientID(fmt.Sprintf("%s-%d", os.Getenv("MQTT_CLIENTID"), rand.Intn(9999)))
	b.config.SetUsername(os.Getenv("MQTT_USERNAME"))
	b.config.SetPassword(os.Getenv("MQTT_PWD"))
	b.config.SetAutoReconnect(true)
}

func (b *MqttBuilder) SetConfig(config *mqtt.ClientOptions) {
	b.config = config
}

func (b *MqttBuilder) SetReceivingChannel(channel chan <- []byte) {
	b.rcvChannel = channel
}

func (b *MqttBuilder) GetMqttHandler() *mqttHandler {
	if b.rcvChannel == nil {
		panic("MQTT Handler - Receiving channel not set")
	}
	return &mqttHandler{
		client:  mqtt.NewClient(b.config),
		rcvCh: b.rcvChannel,
	}
}
