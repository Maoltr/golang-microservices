package mqttClient

import (
	"errors"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"sync"
)

var client mqtt.Client
var once sync.Once

const (
	broker = "tcp://localhost:1883"
	QOS    = 2
	Topic  = "users"
)

func GetClient() (mqtt.Client, error) {
	once.Do(func() {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(broker)
		client = mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			logrus.Warn(token.Error())
		}
	})

	if !client.IsConnected() {
		return nil, errors.New("can not connect to mqttClient broker")
	}

	return client, nil
}
