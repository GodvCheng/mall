package util

import (
	"github.com/go-redis/redis/v8"
	"mall/api/conf"
)

var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.RedisConf.RedisAddr,
		Password: conf.RedisConf.RedisPass,
		DB:       conf.RedisConf.DB,
		PoolSize: conf.RedisConf.PoolSize,
	})
}
