package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, IdleTimeout time.Duration) {

	pool = &redis.Pool {
		MaxIdle: maxIdle, //最大空閒連接數
		MaxActive: maxActive, //數據庫最大連接數
		IdleTimeout: IdleTimeout, //空閒時間
		Dial: func() (redis.Conn, error) { //初始化連接 連接哪個ip
			return redis.Dial("tcp", address)
		},
	}

}