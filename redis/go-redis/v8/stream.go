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

	//XLEN，获取stream中元素数量，也就是消息队列长度
	len, err := rdb.XLen(ctx, "mystreamone").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("XLen: ", len)

	// XRead，从消息队列获取数据，阻塞或非阻塞
	val, err := rdb.XRead(ctx, &redis.XReadArgs{
		Block:   time.Second * 10,               // 如果Block设置为0，表示一直阻塞，默认非阻塞。这里设置阻塞10s
		Count:   2,                              // 读取消息的数量
		Streams: []string{"mystreamone", "0-0"}, // 消息队列名称，从哪个ID开始读起，0-0 表示从mystreamone的第一个ID开始读
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("XRead: ", val)

	// XRANGE，从队列左边获取值，ID 从小到大
	vals, err := rdb.XRange(ctx, "mystreamone", "-", "+").Result() //- + 表示读取所有
	if err != nil {
		panic(err)
	}
	fmt.Println("XRange: ", vals)
	// XRangeN，从队列左边获取N个值，ID 从小到大
	vals, _ = rdb.XRangeN(ctx, "mystreamone", "-", "+", 2).Result() //顺序获取队列前2个值
	fmt.Println("XRangeN: ", vals)

	// XRevRange，从队列右边获取值，ID 从大到小，与XRANGE相反
	vals, _ = rdb.XRevRange(ctx, "mystreamone", "+", "-").Result()
	fmt.Println("XRevRange: ", vals)
	// XRevRangeN，从队列右边获取N个值，ID 从大到小
	// rdb.XRevRangeN(ctx, "mystreamone", "+", "-", 2).Result()

	//XDEL - 删除消息
	//err = rdb.XDel(ctx, "mystreamone", "1678984704869-0").Err()

	// ========= 消费者组相关操作 API ===========

	// XGroupCreate，创建一个消费者组

	/*
		    err = rdb.XGroupCreate(ctx, "mystreamone", "test_group1", "0").Err() // 0-从第一个获取，$-从最新获取
			if err != nil {
				panic(err)
			}
	*/

	// XReadGroup，读取消费者中消息
	readgroupval, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		// Streams第二个参数为ID，list of streams and ids, e.g. stream1 stream2 id1 id2
		// id为 >，表示最新未读消息ID，也是未被分配给其他消费者的最新消息
		// id为 0 或其他，表示可以获取已读但未确认的消息。这种情况下BLOCK和NOACK都会忽略
		// id为具体ID，表示获取这个消费者组的pending的历史消息，而不是新消息
		Streams:  []string{"mystreamone", ">"},
		Group:    "test_group1",    //消费者组名
		Consumer: "test_consumer1", // 消费者名
		Count:    1,
		Block:    0,    // 是否阻塞，=0 表示阻塞且没有超时限制。只要大于1条消息就立即返回
		NoAck:    true, // true-表示读取消息时确认消息
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("XReadGroup: ", readgroupval)

	// XPending，获取待处理的消息
	count, err := rdb.XPending(ctx, "mystreamone", "test_group1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("XPending: ", count)

	// XAck , 将消息标记为已处理
	err = rdb.XAck(ctx, "mystreamone", "test_group1", "1678984704869-0").Err()

	// XClaim ， 转移消息的归属权
	claiminfo, err := rdb.XClaim(ctx, &redis.XClaimArgs{
		Stream:   "mystreamone",
		Group:    "test_group1",
		Consumer: "test_consumer2",
		MinIdle:  time.Second * 10, // 表示要转移的消息需要最少空闲 10s 才能转移
		Messages: []string{"1678984704869-0"},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("XClaim: ", claiminfo)

	// XInfoStream , 获取流的消息
	info, err := rdb.XInfoStream(ctx, "mystreamone").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("XInfoStream: ", info)

	// XInfoGroups , 获取消费者组消息
	groupinfo, _ := rdb.XInfoGroups(ctx, "mystreamone").Result()
	fmt.Println("XInfoGroups: ", groupinfo)

	// XInfoConsumer ，获取消费者信息
	consumerinfo, _ := rdb.XInfoConsumers(ctx, "mystreamone", "test_group1").Result()
	fmt.Println("XInfoConsumers: ", consumerinfo)
}
