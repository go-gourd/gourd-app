package subscribe

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func DeviceBasicHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Device basic message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
