package consumer

import (
	"bitbucket.org/Milinel/golangContainer/models"
	"bitbucket.org/Milinel/golangContainer/services/redisClient"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
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

	oneHourAgo := time.Now().In(time.Local).Truncate(time.Duration(time.Minute))
	if oneHourAgo.After(message.TimeStamp) {
		logrus.Error("Old timestamp")
		return
	}

	cmd := client.ZAdd(redisClient.Channel, redis.Z{Member: messageBytes, Score: float64(message.TimeStamp.In(time.Local).Unix())})
	logrus.Info(cmd.Result())
}

func TTL() {
	go func() {
		ticker := time.NewTicker(time.Minute)

		for range ticker.C {
			client := redisClient.GetClient()
			oneHourAgo := time.Now().In(time.Local).Truncate(time.Duration(time.Hour))
			client.ZRemRangeByScore(redisClient.Channel, "0", strconv.FormatInt(oneHourAgo.Unix(), 10))
		}
	}()
}

func prepareMessage(message models.User) (*models.UserUI, error) {
	if message.FirstName == "" || message.LastName == "" {
		return nil, errors.New("first and last name can't be empty")
	}

	user := message.TransformUser()
	return &user, nil
}
