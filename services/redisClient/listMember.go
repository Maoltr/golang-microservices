package redisClient

import "github.com/go-redis/redis"

// Z represents sorted set member.
type Z struct {
	Score  float64
	Member interface{}
}

func convertRedisZ(z redis.Z) Z {
	return Z{Score: z.Score, Member: z.Member}
}

func convertToRedisZ(z Z) redis.Z {
	return redis.Z{Score: z.Score, Member: z.Member}
}