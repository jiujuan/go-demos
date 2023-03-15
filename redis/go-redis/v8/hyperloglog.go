package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
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

	// 设置hyperloglog的键myset
	for i := 0; i < 10; i++ {
		if err := rdb.PFAdd(ctx, "myset", fmt.Sprint(i)).Err(); err != nil {
			panic(err)
		}
	}

	ctx = context.Background()
	//PFCount, 返回hyperloglog的近似值
	card, err := rdb.PFCount(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("PFCount: ", card)

	// PFMerge，合并2个hyperloglog
	for i := 0; i < 10; i++ {
		if err = rdb.PFAdd(ctx, "myset2", fmt.Sprintf("val%d", i)).Err(); err != nil {
			panic(err)
		}
	}
	rdb.PFMerge(ctx, "mergeset", "myset", "myset2")
	card, _ = rdb.PFCount(ctx, "mergeset").Result()
	fmt.Println("merge: ", card)
}

/*output:
PFCount:  10
merge:  20
*/
