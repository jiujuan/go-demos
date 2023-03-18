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

	ps := rdb.Subscribe(ctx, "mychannel*")
	defer ps.Close()
	// PubSubChannels，查询活跃的频道
	fmt.Println("====PubSubChannels====")
	channels, _ := rdb.PubSubChannels(ctx, "").Result() //"" 为空，查询所有活跃的channel频道
	for ch, v := range channels {
		fmt.Println(ch, v)
	}
	// 指定匹配模式
	channels, _ = rdb.PubSubChannels(ctx, "mychannel*").Result()
	for ch, v := range channels {
		fmt.Println("PubSubChannels* ：", ch, v)
	}

	fmt.Println("====PubSubNumSub====")
	// PubSubNumSub，具体的channel有多少个订阅者
	numsub, _ := rdb.PubSubNumSub(ctx, "mychannel1", "mychannel2").Result()
	for ch, count := range numsub {
		fmt.Println(ch, ",", count) // ch-channel名字，count-channel的订阅者数量
	}

	// PubSubNumPat， 模式匹配
	pubsub := rdb.PSubscribe(ctx, "mychannel*")
	defer pubsub.Close()
	numsubpat, _ := rdb.PubSubNumPat(ctx).Result()
	fmt.Println("PubSubNumPat: ", numsubpat)
}
