package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "",
		DB:          0,
		IdleTimeout: 350,
		PoolSize:    50, // 连接池连接数量
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := rdb.Ping(ctx).Result() // 检查连接redis是否成功
	if err != nil {
		fmt.Println("Connect Failed: %v \n", err)
		panic(err)
	}

	ctx = context.Background()

	// Subscribe，订阅频道
	pubsub := rdb.Subscribe(ctx, "mychannel1", "mychanne2")
	defer pubsub.Close()

	// 退订具体频道
	unsub := pubsub.Unsubscribe(ctx, "mychannel1", "mychannel2")

	// 按照模式匹配退订
	pubsub.PUnsubscribe(ctx, "mychannel*")
}
