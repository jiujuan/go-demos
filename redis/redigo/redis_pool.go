package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     20,
		MaxActive:   120,
		IdleTimeout: 250,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.0.109:6379")
		},
	}
}

func main() {
	client := pool.Get()
	defer client.Close()

	_, err := client.Do("SET", "names", "redis-pool")
	if err != nil {
		fmt.Println("set error: ", err)
		return
	}

	r, err := redis.String(client.Do("GET", "names"))
	if err != nil {
		fmt.Println("get error: ", err)
		return
	}
	fmt.Println(r)
}
