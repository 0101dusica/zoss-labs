package main

import (
	"fmt"
	"log"
	"os"
	"github.com/eclipse/paho.mqtt.golang"
)

func main() {
	broker := "tcp://localhost:1883"
	clientID := "urbanSense-demo-subscriber"
	topic := "sensors/+/data"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to broker: %v", token.Error())
	}

	fmt.Printf("Subscribed to topic: %s\n", topic)

	callback := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("[LOG] Topic: %s | Payload: %s\n", msg.Topic(), msg.Payload())
	}

	if token := client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe: %v", token.Error())
	}

	// Block forever
	select {}
}
