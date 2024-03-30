package subscribe

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SysBrokersClientsHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
