package main

import (
	"github.com/go-redis/redis/v8"
)

// 连接到集群模式
func initCluster() error {
	rdb := redis.NewClusterClient(&redis.ClusterOptons{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004"},
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
