package client

import (
	"strings"

	"github.com/go-redis/redis"
	"github.com/lunarhq/sharedutils/env"
	"github.com/lunarhq/sharedutils/redis/key"
	"github.com/lunarhq/sharedutils/redis/request"
)

var (
	RedisPassword   = env.Get("REDIS_PASSWORD", "")
	RedisMasterName = env.Get("REDIS_MASTER_NAME", "mymaster")
	RedisSentinels  = env.Get("REDIS_SENTINELS", "")
)

type Redis struct {
	Keys     *key.Client
	Requests *request.Client
}

func New() (*Redis, error) {
	rc := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    RedisMasterName,
		SentinelAddrs: strings.Split(RedisSentinels, ","),
		Password:      RedisPassword,
	})

	r := &Redis{
		Keys:     &key.Client{rc},
		Requests: &request.Client{rc},
	}
	return r, nil
}
