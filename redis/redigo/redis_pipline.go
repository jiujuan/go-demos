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

	c.Send("SET", "name1", "redis001")
	c.Send("SET", "name2", "redis002")
	c.Flush()

	v, err := c.Receive()
	fmt.Printf("v: %v, err: %v \n", v, err)

	v, err = c.Receive()
	fmt.Printf("v: %v, err: %v \n", v, err)

	v, err = c.Receive() // 夯住，一直等待
	fmt.Printf("v:%v,err:%v\n", v, err)
}
