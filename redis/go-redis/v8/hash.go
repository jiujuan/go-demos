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
	// HSET，根据key设置field字段值
	err = rdb.HSet(ctx, "hashkey", "field-val", "value-one").Err()
	if err != nil {
		panic(err)
	}
	_ = rdb.HSet(ctx, "hashkey", "field1", "value1", "field2", "value2").Err()
	_ = rdb.HSet(ctx, "hashkey", map[string]interface{}{"field3": "value3", "field4": "value4"}).Err()
	_ = rdb.HSet(ctx, "hashkey-two", []string{"field0", "value0", "field1", "value1"}).Err()

	// HSETNX，如果某个字段不存在则设置值
	ok, err := rdb.HSetNX(ctx, "hashkey", "field1", "oneval").Result() // 字段 field1 已存在，所以返回ok值为false
	if err != nil {
		panic(err)
	}
	fmt.Println("HSetNX bool: ", ok)

	// HGET，根据key和field查询值
	val, err := rdb.HGet(ctx, "hashkey", "field-val").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGet: ", val)

	val, _ = rdb.HGet(ctx, "hashkey-two", "field0").Result()
	fmt.Println("HGet hashkey-two: ", val)

	// HGETALL，获取key的所有field-val值
	fieldvals, err := rdb.HGetAll(ctx, "hashkey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGetAll: ", fieldvals) // 返回 map 类型

	// HMSET，根据hash key设置多个字段值，与上面的 HSet 设置多个值很像
	fieldvalues := make(map[string]interface{})
	fieldvalues["age"] = 23
	fieldvalues["firstname"] = "Chare"
	fieldvalues["lastname"] = "Jimmy"
	err = rdb.HMSet(ctx, "hmsetkey", fieldvalues).Err()
	if err != nil {
		panic(err)
	}
	/*//也可以像上面HSet直接设置map值

	rdb.HMSet(ctx, "hmsetkey", map[string]interface{}{"age":23,"firstname":"Chare","LastName":"Jimmy"}).Err()
	*/

	// HMGET, 根据hash key和多个字段获取值
	vals, err := rdb.HMGet(ctx, "hmsetkey", "age", "lastname").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HMGET vals: ", vals)

	// HEXISTX，某个hashkey中字段field否存在
	ok, _ = rdb.HExists(ctx, "hmsetkey", "lastname").Result()
	fmt.Println("HExists: ", ok) // HExists: true

	// HLen，获取hashkey的字段多少
	len, _ := rdb.HLen(ctx, "hashkey").Result()
	fmt.Println("HLen hashkey： ", len) // HLen hashkey: 5

	// HIncrBy，根据key的field字段的整数值加减一个数值
	age, err := rdb.HIncrBy(ctx, "hmsetkey", "age", -3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HIncrBy : ", age) // HIncrBy :  20

	// HDel，删除字段，支持删除多个字段
	rdb.HSet(ctx, "hashkeydel", map[string]interface{}{"field10": "value10", "field11": "value11", "field12": "value12", "field13": "value13"}).Err()
	rdb.HDel(ctx, "hashkeydel", "field10", "field12") //删除多个字段

	delvals, err := rdb.HGetAll(ctx, "hashkeydel").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGetAll hashkeydel: ", delvals)

}

/**
HSetNX bool:  false
HGet:  value-one
HGet hashkey-two:  value0
HGetAll:  map[field-val:value-one field1:value1 field2:value2 field3:value3 field4:value4]
HMGET vals:  [23 Jimmy]
HExists:  true
HLen hashkey：  5
HIncrBy :  20
HGetAll hashkeydel:  map[field11:value11 field13:value13]
 * */
