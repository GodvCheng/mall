package util

import (
	"context"
	"github.com/redis/go-redis/v9"
	"mall/conf"
)

var Rdb *redis.Client
var Ctx context.Context

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.RedisConf.RedisAddr,
		Password: conf.RedisConf.RedisPass,
		DB:       conf.RedisConf.DB,
		PoolSize: conf.RedisConf.PoolSize,
	})
	Ctx = context.Background()
}
