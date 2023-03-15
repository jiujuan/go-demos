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
	// 设置 key 的值，0 表示永不过期
	err = rdb.Set(ctx, "setkey-1", "value-1", 0).Err()
	if err != nil {
		panic(err)
	}

	// 设置 key 的值的过期时间为 30 秒
	err = rdb.Set(ctx, "setkey-2", "value-2", time.Second*30).Err()
	if err != nil {
		panic(err)
	}

	// 获取key的值
	val, err := rdb.Get(ctx, "setkey-1").Result()
	if err == redis.Nil { // 如果返回 redis.Nil 说明key不存在
		fmt.Println("key not exixt")
	} else if err != nil {
		fmt.Println("Get Val error: ", err)
		panic(err)
	}
	fmt.Println("Get Val: ", val)

	val, _ = rdb.Get(ctx, "setkey-2").Result()
	fmt.Println("Get Val setkey-2: ", val)

	// GetRange，字符串截取操作，返回字符串截取后的值
	val, _ = rdb.GetRange(ctx, "setkey-1", 1, 3).Result()
	fmt.Println("get range: ", val)

	rdb.Set(ctx, "setkey-1", "value-1", time.Second*20).Err()
	rdb.Set(ctx, "setkey-2", "value-2", time.Second*40).Err()

	// Expire, 设置key在某个时间段后过期
	val1, _ := rdb.Expire(ctx, "setkey-1", time.Second*20).Result()
	fmt.Println("expire: ", val1)

	// ExpireAt，设置key在某个时间点后过期
	val2, _ := rdb.ExpireAt(ctx, "setkey-2", time.Now().Add(time.Second*50)).Result()
	fmt.Println("expire at: ", val2)

	// TTL
	expire, err := rdb.TTL(ctx, "setkey-1").Result()
	fmt.Println(expire, err)

	// STRLEN
	strlen, _ := rdb.StrLen(ctx, "setkey-1").Result()
	fmt.Println("strlen: ", strlen)
}
