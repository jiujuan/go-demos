package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// https://github.com/cloudwego/hertz-examples/blob/main/parameter/query/main.go
func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// The request responds to url matching: /welcome?firstname=Jane&lastname=Doe&food=apple&food=fish
	h.GET("/welcome", func(ctx context.Context, c *app.RequestContext) {
		// DefaultQuery() 获取参数值，如果没有获取的还可以给它赋一个默认值
		firstname := c.DefaultQuery("firstname", "Guest")
		// shortcut for c.Request.URL.Query().Get("lastname")
		lastname := c.Query("lastname")

		var favoriteFood []string
		c.QueryArgs().VisitAll(func(key, value []byte) {
			if string(key) == "food" {
				favoriteFood = append(favoriteFood, string(value))
			}
		})
		c.String(consts.StatusOK, "Hello: 'firstname: %s' 'lastname: %s', favorite food: %s", firstname, lastname, favoriteFood)
	})

	// 获取路由参数方法 Param()
	h.GET("/hello/:name", func(ctx context.Context, c *app.RequestContext) {
		name := c.Param("name")
		num := c.Query("number")

		c.JSON(consts.StatusOK, utils.H{
			"name": name,
			"num":  num,
		})

	})

	h.Spin()
}
