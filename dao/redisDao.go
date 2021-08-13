package dao

import "bingo/db"

func Get(key string) (interface{}, error) {
	conn, err := db.GetRedis()
	if err != nil {
		return nil, err
	}
	reply, err := conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Set(key string, args ...interface{}) (interface{}, error) {
	conn, err := db.GetRedis()
	if err != nil {
		return nil, err
	}
	var reply interface{}
	if len(args) > 1 {
		reply, err = conn.Do("SET", args[0], args[1])
	} else {
		reply, err = conn.Do("SET", args[0])
	}
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Mset(key string, field string, value interface{}) (interface{}, error) {
	conn, err := db.GetRedis()
	if err != nil {
		return nil, err
	}
	reply, err := conn.Do("MSET", field, value)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
