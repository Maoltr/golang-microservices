package databaseClient

import "github.com/go-redis/redis"

// SetMember represents sorted set member.
type SetMember struct {
	Score  float64
	Member interface{}
}

func convertRedisZ(z redis.Z) SetMember {
	return SetMember{Score: z.Score, Member: z.Member}
}

func convertToRedisZ(z SetMember) redis.Z {
	return redis.Z{Score: z.Score, Member: z.Member}
}