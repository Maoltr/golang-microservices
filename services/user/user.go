package user

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"bitbucket.org/Milinel/golangContainer/services/natsClient"
	"bitbucket.org/Milinel/golangContainer/services/databaseClient"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"
)

func PushJSON(message []byte) error {
	client, err := natsClient.GetPublisher()
	if err != nil {
		return err
	}

	err = client.Publish(natsClient.Topic, message)
	if err != nil {
		return err
	}

	return nil
}

func GetUsers(period time.Duration) ([]models.UserUI, error) {
	client := databaseClient.GetRangeable()

	messages, err := client.RangeWithScores(databaseClient.Channel, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return sortUsers(period, messages), nil
}

func sortUsers(period time.Duration, messages []databaseClient.SetMember) []models.UserUI {
	var result []models.UserUI
	start := time.Now()

	for i := len(messages) - 1; i >= 0; i-- {
		message := messages[i]

		var user models.UserUI

		err := json.Unmarshal([]byte(message.Member.(string)), &user)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}

		start.In(user.TimeStamp.Location())

		if !(start.Sub(user.TimeStamp) < period) {
			break
		}

		result = append(result, user)
	}

	return result
}
