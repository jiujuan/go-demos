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
	// ZADD，添加一个或多个数据到集合中
	//* 添加一个*/
	n, err := rdb.ZAdd(ctx, "zsetkey", &redis.Z{23.0, "tom"}).Result()
	/* 或把字段写上
	member := &redis.Z{
		Score:  23.0,
		Member: "tom",
	}

	n, err := rdb.ZAdd(ctx, "zsetkey", member).Result()
	if err != nil {
		panic(err)
	}
	*/
	fmt.Println("zadd: ", n)
	val, _ := rdb.ZRange(ctx, "zsetkey", 0, -1).Result()
	fmt.Println("ZRange, zsetkey: ", val)

	//* ZADD批量增加*/
	fruits_price_z := []*redis.Z{
		&redis.Z{Score: 5.0, Member: "apple"},
		&redis.Z{Score: 3.5, Member: "orange"},
		&redis.Z{Score: 6.0, Member: "banana"},
		&redis.Z{Score: 9.1, Member: "peach"},
		&redis.Z{Score: 19.0, Member: "cherry"},
	}
	num, err := rdb.ZAdd(ctx, "fruits_price", fruits_price_z...).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("zadd : ", num)

	// ZRANGE，索引范围返回元素，分数从小到大， 0 到 -1 就是所有元素
	vals, err := rdb.ZRange(ctx, "fruits_price", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ZRange,fruits_price: ", vals)

	// ZREVRANGE，分数从大到小
	vals, err = rdb.ZRevRange(ctx, "fruits_price", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ZRevRange,fruits_price: ", vals)

	// ZRANGEBYSCORE ， offset 和 count 可用于分页
	rangbyscore := &redis.ZRangeBy{
		Min:    "3", // 最小分数
		Max:    "7", // 最大分数
		Offset: 0,   // 开始偏移量
		Count:  4,   // 一次返回多少数据
	}
	vals, err = rdb.ZRangeByScore(ctx, "fruits_price", rangbyscore).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ZRangeByScore: ", vals)

	// ZCOUNT ，统计某个分数内的元素个数
	count, _ := rdb.ZCount(ctx, "fruits_price", "3", "7").Result()
	fmt.Println("ZCount: ", count)

	// ZREVRANGEBYSCOREWITHSCORES, 和 ZRANGEBYSCORE 一样，区别是它不仅返回集合元素，也返回元素对应分数
	rangbyscorewithscores := &redis.ZRangeBy{
		Min:    "3", // 最小分数
		Max:    "7", // 最大分数
		Offset: 0,   // 开始偏移量
		Count:  4,   // 一次返回多少数据
	}
	keyvals, err := rdb.ZRangeByScoreWithScores(ctx, "fruits_price", rangbyscorewithscores).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ZRangeByScoreWithScores: ", keyvals)

	// ZCRORE, 查询集合中元素的分数
	score, _ := rdb.ZScore(ctx, "fruits_price", "peach").Result()
	fmt.Println("ZScore: ", score)

	// ZRANK 根据元素查询在集合中的排名，分数从小到大排序查询
	rank, _ := rdb.ZRank(ctx, "fruits_price", "peach").Result()
	fmt.Println("ZRank: ", rank)

	// ZREM，根据Member删除值，一次可以删除一个或多个
	age_z := []*redis.Z{
		&redis.Z{Score: 20, Member: "tom"},
		&redis.Z{Score: 34, Member: "jim"},
		&redis.Z{Score: 23, Member: "lilei"},
		&redis.Z{Score: 43, Member: "hanxu"},
		&redis.Z{Score: 30, Member: "jimmy"},
		&redis.Z{Score: 55, Member: "MA"},
		&redis.Z{Score: 50, Member: "MB"},
		&redis.Z{Score: 52, Member: "MC"},
		&redis.Z{Score: 54, Member: "MD"},
		&redis.Z{Score: 59, Member: "ME"},
		&redis.Z{Score: 70, Member: "MF"},
		&redis.Z{Score: 75, Member: "MG"},
	}
	rdb.ZAdd(ctx, "people_age", age_z...).Err()

	rdb.ZRem(ctx, "people_age", "jim").Err() // 删除一个
	// rdb.ZRem(ctx, "people_age", "jim", "jimmy").Err() // 删除多个
	agevals, _ := rdb.ZRange(ctx, "people_age", 0, -1).Result()
	fmt.Println("ZRem, ZRange age: ", agevals)

	//ZREMRANGEBYSCORE， 根据分数区间删除
	// rdb.ZRemRangeByScore("people_age", "20", "30").Err()  // 删除 20<=分数<=30
	rdb.ZRemRangeByScore(ctx, "people_age", "20", "(30").Err() // 删除 20<=分数<30

	agevals, _ = rdb.ZRange(ctx, "people_age", 0, -1).Result()
	fmt.Println("ZRemRangeByScore, ZRange age: ", agevals)

	// ZREMRANGEBYRANK，根据分数排名删除
	// 从低分到高分进行排序，然后按照索引删除
	rdb.ZRemRangeByRank(ctx, "people_age", 6, 7) // 低分到高分排序，删除第6个元素到第7个元素
	agevals, _ = rdb.ZRange(ctx, "people_age", 0, -1).Result()
	fmt.Println("ZRemRangeByRank, ZRange age: ", agevals)
	// 如果写成负数，那么从高分开始删除
	// rdb.ZRemRangeByRank(ctx, "people_age", -6, -7)

	// ZIncrBy, 增加分数
	rdb.ZIncrBy(ctx, "people_age", 12, "MG").Err()
	score, _ = rdb.ZScore(ctx, "people_age", "MG").Result()
	fmt.Println("ZScore: ", score)
}

/*
zadd:  0
ZRange, zsetkey:  [tom]
zadd :  0
ZRange,fruits_price:  [orange apple banana peach cherry]
ZRevRange,fruits_price:  [cherry peach banana apple orange]
ZRangeByScore:  [orange apple banana]
ZCount:  3
ZRangeByScoreWithScores:  [{3.5 orange} {5 apple} {6 banana}]
ZScore:  9.1
ZRank:  3
ZRem, ZRange age:  [tom lilei jimmy hanxu MB MC MD MA ME MF MG]
ZRemRangeByScore, ZRange age:  [jimmy hanxu MB MC MD MA ME MF MG]
ZRemRangeByRank, ZRange age:  [jimmy hanxu MB MC MD MA MG]
ZScore:  87
*/
