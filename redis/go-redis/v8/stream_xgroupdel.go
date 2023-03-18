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

	//XGroupDelConsumer，删除消费者
	count, _ := rdb.XGroupDelConsumer(ctx, "mystreamone", "test_group1", "test_consumer1").Result()
	fmt.Println("XGroupDelConsumer: ", count)

	// XGroupDestroy , 删除消费者组
	count, _ = rdb.XGroupDestroy(ctx, "mystreamone", "test_group1").Result()
	fmt.Println("XGroupDestroy: ", count)

}
