package user

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"bitbucket.org/Milinel/golangContainer/services/mqttClient"
	"bitbucket.org/Milinel/golangContainer/services/redisClient"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

func PushJSON(message []byte) error {
	client, err := mqttClient.GetClient()
	if err != nil {
		return err
	}

	token := client.Publish(mqttClient.Topic, mqttClient.QOS, false, message)

	if token.Error() != nil {
		return token.Error()
	}

	if !token.Wait() {
		return errors.New("can not push message into queue")
	}

	return nil
}

func GetUsers(period time.Duration) ([]models.UserUI, error) {
	client := redisClient.GetClient()

	messages, err := client.ZRangeWithScores(redisClient.Channel, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return sortUsers(period, messages), nil
}

func sortUsers(period time.Duration, messages []redis.Z) []models.UserUI {
	var result []models.UserUI
	start := time.Now()

	for i := len(messages) - 1; i >= 0; i-- {
		message := messages[i]

		var user models.UserUI

		err := json.Unmarshal([]byte(message.Member.(string)), &user)
		if err != nil {
			logrus.Error(err.Error())
		}

		start.In(user.TimeStamp.Location())

		if !(start.Sub(user.TimeStamp) < period) {
			break
		}

		result = append(result, user)
	}

	return result
}