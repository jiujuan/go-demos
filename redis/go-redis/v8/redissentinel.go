package main

import (
	"github.com/go-redis/redis/v8"
)

// 连接哨兵模式
func initClient() (err error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptons{
		MasterName:    "master-name",
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
