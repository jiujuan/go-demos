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

	// MSet 设置值
	err = rdb.MSet(ctx, "mset-key1", "mset-val1", "mset-key2", "mset-val2", "mset-key3", "mset-val3").Err()
	if err != nil {
		fmt.Println("MSet ERROR : ", err)
	}
	// MGet 获取值
	vals, err := rdb.MGet(ctx, "mset-key1", "mset-key2", "mset-key3").Result()
	if err != nil {
		fmt.Println("MGet ERROR: ", err)
		panic(err)
	}
	fmt.Println("vals: ", vals)

}
