package consumer

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"bitbucket.org/Milinel/golangContainer/services/databaseClient"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func PushToRedis(user models.User) {
	message, err := prepareMessage(user)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	client := databaseClient.GetAddable()
	cmd := client.Add(databaseClient.Channel, *message)
	logrus.Info(cmd.Err())
}

func TTL() {
	go func() {
		ticker := time.NewTicker(time.Minute)

		for range ticker.C {
			logrus.Info("Tick at: ", time.Now())
			client := databaseClient.GetRemovable()
			oneHourAgo := time.Now().In(time.Local).Add(-time.Duration(time.Hour))
			client.RemRangeByScore(databaseClient.Channel, "0", strconv.FormatInt(oneHourAgo.Unix(), 10))
		}
	}()
}

func prepareMessage(message models.User) (*databaseClient.SetMember, error) {
	if message.FirstName == "" || message.LastName == "" {
		return nil, errors.New("first and last name can't be empty")
	}

	user := message.TransformUser()

	messageBytes, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	oneHourAgo := time.Now().In(time.Local).Add(-time.Duration(time.Hour))
	if oneHourAgo.After(user.TimeStamp) {
		return nil, errors.New("old timestamp")
	}

	return &databaseClient.SetMember{Member: messageBytes, Score: float64(user.TimeStamp.In(time.Local).Unix())}, nil
}
