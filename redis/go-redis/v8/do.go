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

	v := rdb.Do(ctx, "get", "key_does_not_exist").String()
	fmt.Printf("%q \n", v)

	err = rdb.Do(ctx, "set", "set-key", "set-val", "EX", time.Second*120).Err()
	fmt.Println("Do set: ", err)
	v = rdb.Do(ctx, "get", "set-key").String()
	fmt.Println("Do get: ", v)
}
