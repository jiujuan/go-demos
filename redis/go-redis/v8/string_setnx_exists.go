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

	// SetNX， 设置并指定过期时间，仅当 key 不存在时候才设置有效
	err = rdb.SetNX(ctx, "setnx-key", "setnx-val", 0).Err()
	if err != nil {
		fmt.Println("setnx value failed: ", err)
		panic(err)
	}

	// 这里用SetNX设置值，第二次运行后 val2 返回 false，因为第二次运行时 setnx-key2 已经存在
	val2, err := rdb.SetNX(ctx, "setnx-key2", "setnx-val2", time.Second*20).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("val2: %v \n", val2)

	// Exists， 检查某个key是否存在
	n, _ := rdb.Exists(ctx, "setnx-key").Result()
	if n > 0 {
		fmt.Println("n: ", n)
		fmt.Println("set nx key exists")
	} else {
		fmt.Println("set nx key not exists")
	}

	val, _ := rdb.Get(ctx, "setnx-key").Result()
	fmt.Println(val)
}
