package main

import (
	"fmt"
	"net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	go func(topic string) {
		opts := createClientOptions("sub", os.Getenv("CLOUDMQTT_URL"))
		client := mqtt.NewClient(opts)
		client.Start()

		t, _ := mqtt.NewTopicFilter(topic, 0)
		client.StartSubscription(func(client *mqtt.MqttClient, msg mqtt.Message) {
			fmt.Println("Topic=", msg.Topic(), "Payload=", string(msg.Payload()))
		}, t)
	}("#")

	timer := time.NewTicker(1 * time.Second)
	opts := createClientOptions("pub", os.Getenv("CLOUDMQTT_URL"))
	client := mqtt.NewClient(opts)
	client.Start()

	for t := range timer.C {
		client.Publish(0, "currentTime", t.String())
	}
}

func createClientOptions(clientId, raw string) *mqtt.ClientOptions {
	uri, _ := url.Parse(raw)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)

	return opts
}
