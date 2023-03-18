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

	// 一次执行 2 个删除命令
	rdb.Set(ctx, "setkey1", "value1", 0).Err()
	rdb.Set(ctx, "setkey2", "value2", 0).Err()
	//TxPipeline
	txpipe := rdb.TxPipeline()
	txpipe.Del(ctx, "setkey1")
	txpipe.Del(ctx, "setkey2")
	cmds, err := txpipe.Exec(ctx) // 执行 TxPipeline 里的命令
	if err != nil {
		panic(err)
	}
	fmt.Println("TxPipeline: ", cmds)

	// TxPipelined
	var incr2 *redis.IntCmd
	cmds, err = rdb.TxPipelined(ctx, func(txpipe redis.Pipeliner) error {
		txpipe.Set(ctx, "txpipeline_counter2", 30, time.Second*120)
		incr2 = txpipe.Incr(ctx, "txpipeline_counter2")
		txpipe.Expire(ctx, "txpipeline_counter2", time.Second*300)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("TxPipelined: ", incr2.Val())
	fmt.Println("cmds: ", cmds)

}

/*
TxPipeline:  [del setkey1: 1 del setkey2: 1]
TxPipelined:  31
cmds:  [set txpipeline_counter2 30 ex 120: OK incr txpipeline_counter2: 31 expire txpipeline_counter2 300: true]
*/
