package subscribe

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-gourd/gourd/log"
	"strings"
)

type ClientDisconnectStruct struct {
	Username       string `json:"username"`
	Ts             int64  `json:"ts"`
	Sockport       int    `json:"sockport"`
	Reason         string `json:"reason"`
	Protocol       string `json:"protocol"`
	ProtoVer       int    `json:"proto_ver"`
	ProtoName      string `json:"proto_name"`
	Ipaddress      string `json:"ipaddress"`
	DisconnectedAt int64  `json:"disconnected_at"`
	ConnectedAt    int64  `json:"connected_at"`
	Clientid       string `json:"clientid"`
}

type ClientConnectStruct struct {
	Username       string `json:"username"`
	Ts             int64  `json:"ts"`
	Sockport       int    `json:"sockport"`
	Protocol       string `json:"protocol"`
	ProtoVer       int    `json:"proto_ver"`
	ProtoName      string `json:"proto_name"`
	Keepalive      int    `json:"keepalive"`
	Ipaddress      string `json:"ipaddress"`
	ExpiryInterval int    `json:"expiry_interval"`
	ConnectedAt    int64  `json:"connected_at"`
	Clientid       string `json:"clientid"`
	CleanStart     bool   `json:"clean_start"`
}

func SysBrokersClientsHandler(client mqtt.Client, msg mqtt.Message) {

	topic := msg.Topic()

	// 取出topic最后一个/之后的字符串
	topicArr := strings.Split(topic, "/")
	commend := topicArr[len(topicArr)-1]

	if commend == "connected" {
		// 客户端上线
		var clientConnectStruct ClientConnectStruct
		err := json.Unmarshal(msg.Payload(), &clientConnectStruct)
		if err != nil {
			log.Warn("SysBrokersClientsHandler json.Unmarshal err: " + err.Error())
			return
		}

		//TODO: 设备上线

	} else if commend == "disconnected" {
		// 客户端下线
		var clientDisconnectStruct ClientDisconnectStruct
		err := json.Unmarshal(msg.Payload(), &clientDisconnectStruct)
		if err != nil {
			log.Warn("SysBrokersClientsHandler json.Unmarshal err: " + err.Error())
			return
		}

		//TODO: 设备下线

	}

	// 取出消息
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
