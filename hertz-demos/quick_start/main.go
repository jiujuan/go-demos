package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	// Default() 默认地址和端口：http://127.0.0.1:8888

	// 还可以设置地址和端口，server.WithHostPorts("127.0.0.1:8080")，
	// 加入到 Default(server.WithHostPorts("127.0.0.1:8080"))
	h := server.Default()

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.Spin()
}
