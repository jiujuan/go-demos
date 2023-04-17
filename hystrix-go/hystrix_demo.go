package main

import (
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	hystrix.ConfigureCommand("my_command_name", hystrix.CommandConfig{
		Timeout:                1000, // 超时时间1000ms
		MaxConcurrentRequests:  40,   //最大并发数40
		RequestVolumeThreshold: 20,   // 请求数量阙值20，达到这个阙值才可能触发熔断
		ErrorPercentThreshold:  20,   // 错误百分比例阙值 20%
	})

	client := http.Client{}

	doRequest := func() error {
		req, err := http.NewRequest("GET", "https://www.bin.com", nil)
		if err != nil {
			return err
		}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		fmt.Println("doRequest: ", resp.Status)

		return nil
	}

	err := hystrix.Do("my_command_name", func() error {
		return doRequest()
	}, nil)

	if err != nil {
		fmt.Println("end: ", err)
	}
}
