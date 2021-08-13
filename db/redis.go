package db

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

func InitRedis() {
	redisPool = &redis.Pool{
		MaxIdle:     16,  // 最初的连接数量
		MaxActive:   100, //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭)
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func getRedis() (redis.Conn, error) {
	conn := redisPool.Get()
	if _, ok := conn.(redis.Conn); !ok {
		return nil, errors.New("获取连接失败")
	}
	return conn, nil
}
