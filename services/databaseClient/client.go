package databaseClient

import "github.com/go-redis/redis"

type Rangeable interface {
	RangeWithScores(key string, start, stop int64) *Cmd
}

type Addable interface {
	Add(key string, members ...SetMember) *Cmd
}

type Removable interface {
	RemRangeByScore(key, min, max string) *Cmd
}

type Client struct {
	redisClient *redis.Client
}

func (c *Client) RangeWithScores(key string, start, stop int64) *Cmd {
	cmd := c.redisClient.ZRangeWithScores(key, start, stop)
	redisVal, err := cmd.Result()

	var val []SetMember

	for _, v := range redisVal {
		val = append(val, convertRedisZ(v))
	}

	return &Cmd{val: val, err: err}
}

func (c *Client) Add(key string, members ...SetMember) *Cmd {
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
