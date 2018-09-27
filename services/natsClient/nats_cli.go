package natsClient

import (
	"errors"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"sync"
)

var client *nats.Conn
var once sync.Once
var err error

const (
	QOS   = 2
	Topic = "users"
)

type Subscriber interface {
	Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error)
}

type Publisher interface {
	Publish(subj string, data []byte) error
}

func GetPublisher() (Publisher, error) {
	once.Do(func() {
		client, err = nats.Connect(nats.DefaultURL)
	})

	if err != nil || !client.IsConnected() {
		logrus.Fatal(err)
		return nil, errors.New("no connection to nats")
	}

	return client, nil
}

func GetSubscriber() (Subscriber, error) {
	once.Do(func() {
		client, err = nats.Connect(nats.DefaultURL)
	})

	if err != nil || !client.IsConnected() {
		logrus.Fatal(err)
		return nil, errors.New("no connection to nats")
	}

	return client, nil
}
