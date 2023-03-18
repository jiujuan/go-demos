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

	// Subscribe，订阅频道接收消息
	pubsub := rdb.PSubscribe(ctx, "mychannel*")
	defer pubsub.Close()

	// 第一种接收消息方法
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

	// 第二种接收消息方法
	// for {
	// 	msg, err := pubsub.ReceiveMessage(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(msg.Channel, msg.Payload)
	// }

}
