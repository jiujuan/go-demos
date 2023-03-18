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
	// GEOADD，添加一个
	val, err := rdb.GeoAdd(ctx, "town-geo-key", &redis.GeoLocation{
		Longitude: 113.2442,
		Latitude:  23.12592,
		Name:      "niwan-town",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("GeoAdd: ", val)
	// GEOADD，添加多个
	val, _ = rdb.GeoAdd(ctx, "town-geo-key",
		&redis.GeoLocation{Longitude: 113.2442, Latitude: 23.12592, Name: "niwan-town"},
		&redis.GeoLocation{Longitude: 113.38397, Latitude: 22.93599, Name: "panyu-town"},
		&redis.GeoLocation{Longitude: 113.60845, Latitude: 22.77144, Name: "nansha-town"},
		&redis.GeoLocation{Longitude: 113.829579, Latitude: 23.290497, Name: "zengcheng-town"},
	).Result()
	fmt.Println("Mulit GeoAdd : ", val)

	// GEOPOS，根据名字获取经纬度
	lonlats, err := rdb.GeoPos(ctx, "town-geo-key", "zengcheng-town", "panyu-town").Result()
	if err != nil {
		panic(err)
	}
	for _, lonlat := range lonlats {
		fmt.Println("GeoPos, ", "Longitude: ", lonlat.Longitude, "Latitude: ", lonlat.Latitude)
	}

	// GEODIST , 计算两地距离
	distance, err := rdb.GeoDist(ctx, "town-geo-key", "niwan-town", "nansha-town", "m").Result() // m-米，km-千米，mi-英里
	if err != nil {
		panic(err)
	}
	fmt.Println("GeoDist: ", distance, " m")

	// GEOHASH，计算hash值
	hash, _ := rdb.GeoHash(ctx, "town-geo-key", "zengcheng-town").Result()
	fmt.Println("zengcheng-town geohash: ", hash)

	// GEORADIUS，计算范围内包含的经纬度位置
	radius, _ := rdb.GeoRadius(ctx, "town-geo-key", 113.829579, 23.290497, &redis.GeoRadiusQuery{
		Radius:      800,
		Unit:        "km",
		WithCoord:   true,  // WITHCOORD参数，返回结果会带上匹配位置的经纬度
		WithDist:    true,  // WITHDIST参数，返回结果会带上匹配位置与给定地理位置的距离。
		WithGeoHash: true,  // WITHHASH参数，返回结果会带上匹配位置的hash值。
		Count:       4,     // COUNT参数，可以返回指定数量的结果。
		Sort:        "ASC", // 传入ASC为从近到远排序，传入DESC为从远到近排序。
	}).Result()
	for _, v := range radius {
		fmt.Println("GeoRadius: ", v)
	}
	// 上面式子里参数更多详情请看这里：http://redisdoc.com/geo/georadius.html

}

/*
GeoAdd:  0
Mulit GeoAdd :  0
GeoPos,  Longitude:  113.8295790553093 Latitude:  23.290497021802757
GeoPos,  Longitude:  113.3839675784111 Latitude:  22.935990920457606
GeoDist:  54280.9773  m
zengcheng-town geohash:  [ws0uqrbhvr0]
GeoRadius:  {zengcheng-town 113.8295790553093 23.290497021802757 0 4046592114973855}
GeoRadius:  {panyu-town 113.3839675784111 22.935990920457606 60.2724 4046531372960175}
*/
