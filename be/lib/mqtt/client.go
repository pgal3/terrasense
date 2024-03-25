package mqtt_lib

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttOptions struct {
	Host     string
	Port     string
	ClientId string
	Username string
	Password string
}

type MqttClient struct {
	Client  mqtt.Client
}

func New(options MqttOptions) *MqttClient {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%s", options.Host, options.Port))
	opts.ClientID = options.ClientId
	opts.Username = options.Username
	opts.Password = options.Password
	c := &MqttClient{
		Client: mqtt.NewClient(opts),
	}
	if token := c.Client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
	return c
}

func (c *MqttClient) Pub(topic string, payload any) error {
	token := c.Client.Publish(topic, 0, true, payload)
	token.Wait()
	if token.Error() != nil {
		fmt.Println("Failed to publish mqtt message")
		fmt.Println(token.Error().Error())
		return token.Error()
	}
	return nil
}

func (c *MqttClient) Sub(topic string, handler func([]byte)){
	mqttHandler := func(cl mqtt.Client, m mqtt.Message){
		fmt.Printf("Message received from topic %s with messageId %d",m.Topic(), m.MessageID())
		handler(m.Payload())
	}
	token := c.Client.Subscribe(topic, 0, mqttHandler)
	token.Wait()
}
