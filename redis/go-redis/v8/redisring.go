package main

import (
	"github.com/go-redis/redis/v8"
)

//https://redis.uptrace.dev/
func main() {
	rdb := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"shared1": ":7000",
			"shared2": ":7001",
			"shared3": ":7002",
		},
	})

	if err := rdb.Set(ctx, "foo", "bar").Err(); err != nil {
		panic(err)
	}
}
