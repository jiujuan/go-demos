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

	// 向频道 mychannel1 发布消息 payload1
	err = rdb.Publish(ctx, "mychannel1", "payload1").Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Publish(ctx, "mychannel1", "hello").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	rdb.Publish(ctx, "mychannel2", "hello2").Err()
}
