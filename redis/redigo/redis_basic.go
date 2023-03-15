package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.0.109:6379")
	if err != nil {
		fmt.Println("conn redis failed, err: ", err)
		return
	}

	defer c.Close()

	//set
	_, err = c.Do("SET", "name", "redis-go")
	if err != nil {
		fmt.Println("err")
		return
	}
	//get
	r, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

	//hset
	_, err = c.Do("HSET", "names", "redis", "hset")
	if err != nil {
		fmt.Println(err)
		return
	}
	//hget
	r, err = redis.String(c.Do("HGET", "names", "redis"))
	if err != nil {
		fmt.Println("hget err: ", err)
		return
	}
	fmt.Println(r)

	//exipres
	_, err = c.Do("expires", "names", 5)
	if err != nil {
		fmt.Println("expire err: ", err)
		return
	}

}
