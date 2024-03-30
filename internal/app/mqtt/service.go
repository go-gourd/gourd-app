package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-gourd/gourd/config"
	"gourd/internal/app/mqtt/subscribe"
)

// Config 配置
type Config struct {
	Broker   string `toml:"broker" json:"broker"`
	Port     int    `toml:"port" json:"port"`
	Username string `toml:"username" json:"username"`
	Password string `toml:"password" json:"password"`
	ClientId string `toml:"client_id" json:"client_id"`
}

// ServiceStart 启动mqtt服务
func ServiceStart() {

	mqttConfig := &Config{}
	err := config.Unmarshal("mqtt", mqttConfig)
	if err != nil {
		panic(err)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttConfig.Broker, mqttConfig.Port))
	opts.SetUsername(mqttConfig.Username)
	opts.SetPassword(mqttConfig.Password)
	opts.SetClientID(mqttConfig.ClientId)
	opts.SetDefaultPublishHandler(subscribe.DefaultHandler)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅客户端事件
	topic := "$SYS/brokers/+/clients/#"
	token := client.Subscribe(topic, 1, subscribe.SysBrokersClientsHandler)
	token.Wait()

	select {}
}
