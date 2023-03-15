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

	rdb.Set(ctx, "setkey-expire-1", "value-expire-1", 0).Err()
	rdb.Set(ctx, "setkey-expire-2", "value-expire-2", time.Second*40).Err()

	// Expire, 设置key在某个时间段后过期
	val1, _ := rdb.Expire(ctx, "setkey-expire-1", time.Second*20).Result()
	fmt.Println("expire: ", val1)

	// ExpireAt，设置key在某个时间点后过期
	val2, _ := rdb.ExpireAt(ctx, "setkey-expire-2", time.Now().Add(time.Second*50)).Result()
	fmt.Println("expire at: ", val2)

	// TTL
	expire, err := rdb.TTL(ctx, "setkey-expire-1").Result()
	fmt.Println(expire, err)
}
