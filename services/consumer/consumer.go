package consumer

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"bitbucket.org/Milinel/golangContainer/services/redisClient"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func PushToRedis(user models.User) {
	message, err := prepareMessage(user)
	if err != nil {
		logrus.Error(err.Error())
	}

	client := redisClient.GetClient()

	messageBytes, err := json.Marshal(message)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	client.ZAdd(redisClient.Channel, redis.Z{Member: messageBytes, Score: float64(message.TimeStamp.Unix())})
}

func prepareMessage(message models.User) (*models.UserUI, error) {
	if message.FirstName == "" || message.LastName == "" {
		return nil, errors.New("first and last name can't be empty")
	}

	user := message.TransformUser()
	return &user, nil
}
