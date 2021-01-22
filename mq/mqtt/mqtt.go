package mqtt

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/yosssi/gmq/mqtt"
	MQTTClient "github.com/yosssi/gmq/mqtt/client"
	"time"
)

/**
 * @author Emove
 * @date 2021/1/6
 */

const topic = "test-topic"

type MqttClient struct {
	Timeout  int
	UserName string
	Password string
	Addr     string
	Client   *MQTTClient.Client
}

type MqttPubBody struct {
	Client  string `json:"client_id"`
	Qos     int    `json:"qos"`
	Retain  bool   `json:"retain"`
	Topic   string `json:"topic"`
	Payload string `json:"payload"`
}

func NewTcpMqttClient() (*MqttClient, error) {
	client := MQTTClient.New(&MQTTClient.Options{
		ErrorHandler: func(e error) {
			logrus.Errorf("mqtt tcp link errors: %s", e.Error())
		},
	})
	err := client.Connect(&MQTTClient.ConnectOptions{
		Network:      "tcp",
		Address:      "192.168.246.129:1883",
		ClientID:     []byte("MqttTestClient"),
		CleanSession: true,
	})
	if err != nil {
		return nil, err
	}
	return &MqttClient{
		Addr:     "192.168.246.129:1883",
		Timeout:  3,
		UserName: "admin",
		Password: "public",
		Client:   client,
	}, nil
}

func (mc *MqttClient) Publish(message string) error {
	if mc.Client == nil {
		return errors.Errorf("mc client not init, please use NewTcpMqttClient init")
	}
	err := mc.Client.Publish(&MQTTClient.PublishOptions{
		QoS:       mqtt.QoS1,
		Retain:    true,
		TopicName: []byte(topic),
		Message:   []byte(message),
	})
	if err != nil {
		return err
	}
	return nil
}

func (mc *MqttClient) Subscribe() {
	client := mc.Client
	var subscribeHandlers []*MQTTClient.SubReq
	subscribeHandlers = append(subscribeHandlers, &MQTTClient.SubReq{
		TopicFilter: []byte(topic),
		QoS:         mqtt.QoS1,
		Handler: func(topicName, message []byte) {
			fmt.Printf("receive: topci: %s, message: %s \n", string(topicName), string(message))
		},
	})
	// sub
	_ = client.Subscribe(&MQTTClient.SubscribeOptions{
		SubReqs: subscribeHandlers,
	})
}

func TestMqtt() {

	mc, createErr := NewTcpMqttClient()
	if createErr != nil {
		logrus.Panicf("mqtt tcp link errors: %s", createErr.Error())
	}

	cli := mc.Client
	// Terminate the Client.
	defer cli.Terminate()

	// Subscribe to topics.
	mc.Subscribe()

	for i := 0; i < 20; i++ {
		_ = mc.Publish(fmt.Sprintf("test message %d", i))
		time.Sleep(1 * time.Millisecond)
	}

	time.Sleep(120 * time.Second)

	// Disconnect the Network Connection.
	if err := cli.Disconnect(); err != nil {
		panic(err)
	}
}
