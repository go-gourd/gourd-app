package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-gourd/gourd/config"
	"gourd/internal/app/mqtt/subscribe"
	"strconv"
	"time"
)

// Config 配置
type Config struct {
	Broker      string `toml:"broker" json:"broker"`
	Username    string `toml:"username" json:"username"`
	Password    string `toml:"password" json:"password"`
	ClientId    string `toml:"client_id" json:"client_id"`
	SharePrefix string `toml:"share_prefix" json:"share_prefix"`
	Distributed bool   `toml:"distributed" json:"distributed"`
}

// ServiceStart 启动mqtt服务
func ServiceStart() {

	mqttConfig := &Config{}
	err := config.Unmarshal("mqtt", mqttConfig)
	if err != nil {
		panic(err)
	}

	// 如果需分布式部署，则随机生成一个客户端id
	if mqttConfig.Distributed {
		// 生成当前时间纳秒的16进制作为当前客户端id标识
		nanoBytes := strconv.FormatInt(time.Now().UnixNano(), 16)
		mqttConfig.ClientId += nanoBytes
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttConfig.Broker)
	opts.SetUsername(mqttConfig.Username)
	opts.SetPassword(mqttConfig.Password)
	opts.SetClientID(mqttConfig.ClientId)
	opts.SetDefaultPublishHandler(subscribe.DefaultHandler)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅客户端事件
	topic := mqttConfig.SharePrefix + "$SYS/brokers/+/clients/#"
	token := client.Subscribe(topic, 1, subscribe.SysBrokersClientsHandler)
	token.Wait()

	// 订阅客户端事件
	topic = mqttConfig.SharePrefix + "device/basic/#"
	token = client.Subscribe(topic, 1, subscribe.DeviceBasicHandler)
	token.Wait()

	select {}
}
