package databaseClient

import (
	"github.com/go-redis/redis"
	"sync"
)

const (
	address  = "localhost:6379"
	password = ""
	db       = 0
	Channel  = "users"
)

var redisClient *redis.Client
var client *Client
var once sync.Once

func GetRangeable() Rangeable {
	return getClient()
}

func GetAddable() Addable {
	return getClient()
}

func GetRemovable() Removable {
	return getClient()
}

func getClient() *Client {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		})

		client = &Client{redisClient: redisClient}
	})

	return client
}
