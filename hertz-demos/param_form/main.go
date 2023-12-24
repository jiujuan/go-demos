package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// content-type : application/x-www-form-urlencoded
	h.POST("/urlencode", func(ctx context.Context, c *app.RequestContext) {
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.PostArgs().VisitAll(func(key, value []byte) {
			if string(key) == "name" {
				fmt.Printf("This is %s!", string(value))
			}
		})

		c.JSON(consts.StatusOK, utils.H{
			"name":    name,
			"message": message,
		})
	})

	// content-type : multipart/form-data
	h.POST("/form-data", func(ctx context.Context, c *app.RequestContext) {
		id := c.FormValue("id")
		name := c.FormValue("name")
		message := c.FormValue("message")

		c.String(consts.StatusOK, "id: %s; name: %s; message: %s\n", id, name, message)
	})

	h.Spin()
}
