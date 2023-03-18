package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
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

	var incr func(string) error
	incr = func(key string) error {
		err = rdb.Watch(ctx, func(tx *redis.Tx) error { //Watch 监控函数
			n, err := tx.Get(ctx, key).Int64() // 先查询下当前watch监听的key的值
			if err != nil && err != redis.Nil {
				return err
			}

			// 如果key的值没有改变的话，pipe 函数才会调用成功
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, key, strconv.FormatInt(n+1, 10), 0)
				return nil
			})
			return err
		}, key)

		if err == redis.TxFailedErr {
			return incr(key)
		}
		return err
	}

	keyname := "keynameone"
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := incr(keyname)
			fmt.Println("[for] err: ", err)
		}()
	}
	wg.Wait()

	n, err := rdb.Get(ctx, keyname).Int64()
	if err != nil {
		panic(err)
	}
	fmt.Println("last key val: ", n)
}
