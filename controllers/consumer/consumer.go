package consumer

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"bitbucket.org/Milinel/golangContainer/services/consumer"
	"bitbucket.org/Milinel/golangContainer/services/natsClient"
	"encoding/json"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
)

func Listen() {
	var messageChan chan models.User
	messageChan = make(chan models.User)

	client, err := natsClient.GetClient()
	if err != nil {
		panic(err)
	}

	client.Subscribe(natsClient.Topic, func(msg *nats.Msg) {
		var user models.User

		err := json.Unmarshal(msg.Data, &user)
		if err != nil {
			logrus.Error(err.Error())
		}

		messageChan <- user
	})
	consumer.TTL()
	logrus.Infof("Subscribed to topic: %s", natsClient.Topic)

	for {
		message := <-messageChan
		consumer.PushToRedis(message)
	}
}
