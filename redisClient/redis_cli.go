package redisClient

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
)

const (
	address = "localhost:6379"
	password = ""
	db = 0
	Channel = "users"
)

var client *redis.Client
var once sync.Once

func GetClient() *redis.Client{
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr: address,
			Password: password,
			DB: db,
		})
		client.Expire(Channel, time.Duration(time.Hour))
	})

	return client
}
