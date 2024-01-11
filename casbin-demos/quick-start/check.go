package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("Enforcer Failed: %v\n", err)
	}

	sub := "alice" // 用户
	obj := "data1" // 资源
	act := "read"  // 用户想在资源上的操作权限

	ok, err := e.Enforce(sub, obj, act)
	fmt.Println(ok)
	if ok {
		fmt.Println("ACCESS")
	} else {
		fmt.Println("Deny ", err)
	}

	fmt.Println("============")
	act = "write"
	ok, err = e.Enforce(sub, obj, act)
	fmt.Println(ok)
	if ok {
		fmt.Println("ACCESS")
	} else {
		fmt.Printf("Deny %v", err)
	}
}
