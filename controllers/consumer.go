package controllers

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"bitbucket.org/Milinel/golangContainer/mqttClient"
	"bitbucket.org/Milinel/golangContainer/services"
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

func Listen() {
	var messageChan chan models.User
	messageChan = make(chan models.User)

	client, err := mqttClient.GetClient()
	if err != nil {
		panic(err)
	}

	client.AddRoute(mqttClient.Topic, func(client mqtt.Client, message mqtt.Message) {
		var user models.User

		err := json.Unmarshal(message.Payload(), &user)
		if err != nil {
			logrus.Error(err.Error())
		}

		messageChan <- user
	})

	if token := client.Subscribe(mqttClient.Topic, mqttClient.QOS, nil); token.Wait() && token.Error() != nil {
		logrus.Fatal(token.Error())
	}

	logrus.Infof("Subscribed to topic: %s", mqttClient.Topic)

	for {
		message := <-messageChan
		services.PushToRedis(message)
	}
}
