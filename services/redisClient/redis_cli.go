package redisClient

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

type Client struct {
	redisClient *redis.Client
}

func (c *Client) RangeWithScores(key string, start, stop int64) *Cmd {
	cmd := c.redisClient.ZRangeWithScores(key, start, stop)
	redisVal, err := cmd.Result()

	var val []Z

	for _, v := range redisVal {
		val = append(val, convertRedisZ(v))
	}

	return &Cmd{val: val, err: err}
}

func (c *Client) Add(key string, members ...Z) *Cmd {
	var val []redis.Z

	for _, v := range members {
		val = append(val, convertToRedisZ(v))
	}

	cmd := c.redisClient.ZAdd(key, val...)
	return &Cmd{err: cmd.Err()}
}

func (c *Client) RemRangeByScore(key, min, max string) *Cmd {
	cmd := c.redisClient.ZRemRangeByScore(key, min, max)
	return &Cmd{err: cmd.Err()}
}


type Rangeable interface {
	RangeWithScores(key string, start, stop int64) *Cmd
}

type Addable interface {
	Add(key string, members ...Z) *Cmd
}

type Removable interface {
	RemRangeByScore(key, min, max string) *Cmd
}

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
