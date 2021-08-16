package db

import "github.com/gomodule/redigo/redis"

type RedisPoolInterface interface {
	Get() redis.Conn
	Close() error
}

type RedisType int

type Config struct {
}

const (
	REDIS = iota
	REDIS_SENTINEL
	REDIS_CLUSTER
	CODIS
)

var redisIplPool RedisPoolInterface // redis实例

// InitRedis 根据传入参数不同，初始化不同的redis集群
func InitRedis(redisType RedisType, config Config) {
	switch redisType {
	case REDIS:
		redisIplPool = initRedis()
	case REDIS_SENTINEL:

	case REDIS_CLUSTER:
	case CODIS:
	}
}
