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

	pipe := rdb.Pipeline()
	pipe.Del(ctx, "setkey1")
	pipe.Del(ctx, "setkey2")
	cmds, err := pipe.Exec(ctx) // 执行 pipeline
	if err != nil {
		panic(err)
	}
	fmt.Println("Pipeline: ", cmds)

	// 一次执行 写和过期时间命令，用 pipeline 一次执行这2条命令
	incr := pipe.Incr(ctx, "pipeline_counter")           // Incr 相当于写入
	pipe.Expire(ctx, "pipeline_counter", time.Second*60) // 加上过期时间

	cmds, err = pipe.Exec(ctx) // 执行 pipeline 里的命令
	if err != nil {
		panic(err)
	}
	// 执行 pipe.Exec() 后获取结果
	fmt.Println("Pipeline: ", incr.Val())

	// Pipelined, 另外一种方法 Pipelined
	var incr2 *redis.IntCmd
	cmds, err = rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr2 = pipe.Incr(ctx, "pipeline_counter2")
		pipe.Expire(ctx, "pipeline_counter2", time.Second*60)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Pipelined: ", incr2.Val())

	// Pipelined, 遍历 pipeline 命令返回值
	cmds, err = rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 5; i++ {
			pipe.Set(ctx, fmt.Sprintf("key%d", i), fmt.Sprintf("val%d", i), 0)
			// pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.StatusCmd).Val())
	}

}
