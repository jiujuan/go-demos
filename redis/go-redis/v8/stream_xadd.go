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

	// XADD，添加消息到对尾（这个代码每运行一次就增加一次内容）
	err = rdb.XAdd(ctx, &redis.XAddArgs{
		Stream:     "mystreamone", // 设置流stream的 key，消息队列名
		NoMkStream: false,         //为false，key不存在会新建
		MaxLen:     10000,         //消息队列最大长度，队列长度超过设置最大长度后，旧消息会被删除
		Approx:     false,         //默认false，设为true时，模糊指定stram的长度
		ID:         "*",           //消息ID，* 表示由Redis自动生成
		Values: []interface{}{ //消息队列的内容，键值对形式
			"apple", "12.0",
			"orange", "5.6",
			"banana", "7.6",
		},
		// MinID: "id",//超过设置长度值，丢弃小于MinID消息id
		// Limit: 1000, //限制长度，基本不用
	}).Err()
	if err != nil {
		panic(err)
	}
}
