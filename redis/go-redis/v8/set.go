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
	// SADD，将一个或多个元素数据添加到集合中
	err = rdb.SAdd(ctx, "setkey:1", 20, "dog").Err()
	if err != nil {
		panic(err)
	}
	rdb.SAdd(ctx, "setkey:1", []string{"hanmeimei", "lilei", "tom", "dog", "one"}) // 切片增加数据，dog只有一个数据
	rdb.SAdd(ctx, "setkey:2", []string{"jimmy", "pig", "dog", "lilei"})

	// SMEMBERS，获取集合中的所有元素数据
	smembers, err := rdb.SMembers(ctx, "setkey:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("SMembers，setkey:1: ", smembers)

	// SCARD，获取集合中的元素数量
	scards, err := rdb.SCard(ctx, "setkey:2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("SCard,setkey:2: ", scards)

	// SPOP，随机移除一个数据并返回这个数据
	rdb.SAdd(ctx, "setkey:3", []string{"one", "two", "three", "four", "six"})
	spop, _ := rdb.SPop(ctx, "setkey:3").Result()
	res, _ := rdb.SMembers(ctx, "setkey:3").Result()
	fmt.Println("spop: ", spop, ", SMembers: ", res)
	// SPOPN，随机移除多个元素并返回
	spopn, _ := rdb.SPopN(ctx, "setkey:3", 2).Result()
	res, _ = rdb.SMembers(ctx, "setkey:3").Result()
	fmt.Println("spopn: ", spopn, ", SMembers: ", res)

	// SISMEMBER，判断元素是否在集合中
	ok, err := rdb.SIsMember(ctx, "setkey:3", "two").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("SIsMember, two : ", ok)

	// SDIFF,差集，SDIFF key1,key2 与 SDIFF key1,key2 差集是不同，看下面的例子
	diff, _ := rdb.SDiff(ctx, "setkey:1", "setkey:2").Result()
	fmt.Println("sdiff: ", diff)
	diff2, _ := rdb.SDiff(ctx, "setkey:2", "setkey:1").Result()
	fmt.Println("sdiff2: ", diff2)
	// SUNION,并集
	union, _ := rdb.SUnion(ctx, "setkey:1", "setkey:2").Result()
	fmt.Println("union: ", union)
	// SINTER,交集
	inter, _ := rdb.SInter(ctx, "setkey:1", "setkey:2").Result()
	fmt.Println("inter: ", inter)

	// SREM , 删除值，返回删除元素个数
	rdb.SAdd(ctx, "setkey:4", []string{"one", "two", "three"})
	count, err := rdb.SRem(ctx, "setkey:4", "one", "three").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("SRem: ", count)
}

/*
SMembers，setkey:1:  [20 hanmeimei one lilei tom dog]
SCard,setkey:2:  4
spop:  six , SMembers:  [four three one two]
spopn:  [one two] , SMembers:  [four three]
SIsMember, two :  false
sdiff:  [tom 20 hanmeimei one]
sdiff2:  [jimmy pig]
union:  [hanmeimei one jimmy lilei tom dog 20 pig]
inter:  [lilei dog]
SRem:  2
*/
