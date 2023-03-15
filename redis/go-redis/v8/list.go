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
	// LPUSH 从头部(左边)插入数据，最后的值在最前面
	count, err := rdb.LPush(ctx, "listkeyone", "one", "two", "three", "four").Result()
	if err != nil {
		fmt.Println("lpush err：", err)
	}
	fmt.Println("lpush count: ", count)

	// LRANGE 返回列表范围数据。例子中返回 0 到 -1，就是返回所有数据
	rangeval, err := rdb.LRange(ctx, "listkeyone", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("LRange values: ", rangeval)

	// LLen 返回列表数据大小
	len, err := rdb.LLen(ctx, "listkeyone").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("llen: ", len)

	// LInsert 在指定位置插入数据
	err = rdb.LInsert(ctx, "listkeyone", "before", "two", 2).Err()
	if err != nil {
		panic(err)
	}

	vals, _ := rdb.LRange(ctx, "listkeyone", 0, -1).Result()
	fmt.Println("LInsert val: ", vals)

	// RPUSH 在 list 尾部插入值
	count, err := rdb.RPush(ctx, "listkeyone", "six", "five").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("RPush count: ", count)

	// RPOP 删除list列表尾部(右边)值
	val, err := rdb.RPop(ctx, "listkeyone").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("rpop val: ", val)
	vals, _ = rdb.LRange(ctx, "listkeyone", 0, -1).Result()
	fmt.Println("(rpop)lrange val: ", vals)

	// LPOP 删除list列表头部(左边)值
	val, err = rdb.LPop(ctx, "listkeyone").Result()
	fmt.Println("rpop val: ", val)

	// LIndex 根据索引查询值，索引是从0开始
	val1, _ := rdb.LIndex(ctx, "listkeyone", 3).Result()
	fmt.Println("LIndex val: ", val1)

	// LSET 根据索引设置某个值，索引从0开始
	val2, _ := rdb.LSet(ctx, "listkeyone", 3, "han").Result()
	fmt.Println("lset: ", val2)

	// LREM 删除列表中的数据
	del, err := rdb.LRem(ctx, "listkeyone", 1, 5) // 从列表左边开始删除值 5，出现重复元素只删除一次
	if err != nil {
		panic(err)
	}
	fmt.Println("del : ", del)

	rdb.LRem(ctx, "listkeyone", 2, 5) // 从列表头部(左边)开始删除值 5，如果存在多个值 5，则删除 2 个 5

	rdb.LRem(ctx, "listkeyone", -3, 6) // 从列表尾部(右边)开始删除值 6，如果存在多个值 6， 则删除 3 个 6

}
