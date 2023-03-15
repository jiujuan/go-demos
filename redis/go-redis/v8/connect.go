package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
    "context"
)

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password:""，
        DB:0,
    })

    ctx := context.Background()
    val, err := rdb.Get(ctx, "key").Result()
    switch {
    case err == redis.Nil:
        fmt.Println("key does not exist")
    case err!=nil:
        fmt.Println("Get failed", err)
    case val == "":
        fmt.Println("value is empty")
    }


    fmt.Println(val)

    // or
    // get := rdb.Get(ctx, "key")
    // get.Val()
    // get.Err()
}

func Do(rdb *redis.Client) {
    ctx := context.Background()
    val, err := rdb.Do(ctx, "get", "key").Result
    if err != nil {
        if err == redis.Nil {
            fmt.Println("key does not exists")
            return
        }
        panic(err)
    }
    fmt.Println(val.(string))
}