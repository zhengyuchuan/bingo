package db

import (
	"github.com/gomodule/redigo/redis"
)

type RedisPoolInterface interface {
	Get() redis.Conn
	Close() error
}

type RedisType int

type Config struct {
	addr map[string]string
}

const (
	REDIS RedisType = iota
	REDIS_SENTINEL
	REDIS_CLUSTER
	CODIS
)

var redisIplPool RedisPoolInterface // redis实例

// InitRedis 根据传入参数不同，初始化不同的redis集群
func InitRedis(redisType RedisType) {

	switch redisType {
	case REDIS:
		redisIplPool = initRedis()
	case REDIS_SENTINEL:
		redisIplPool = initSentinel()
	case REDIS_CLUSTER:
		redisIplPool = initCluster()
	case CODIS:
		redisIplPool = initCodis()
	}
}
