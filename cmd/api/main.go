package main

import (
	// "context"
	// "errors"
	// "fmt"

	"github.com/AgSword/simpleDouyin/cmd/api/handlers"
	// "github.com/Liujony/tiktokapi/cmd/api/rpc"

	// "github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

var (
	ServiceAddr = "0.0.0.0:8081"
)

// 初始化 Hertz
func InitHertz() *server.Hertz {
	h := server.Default(server.WithHostPorts(ServiceAddr))
	return h
}

// 注册 Router组
func registerGroup(h *server.Hertz) {
	douyin := h.Group("/douyin")

	favorite := douyin.Group("/favorite")
	favorite.GET("/action/", handlers.FavoriteAction)

	user := douyin.Group("/user")
	user.GET("/login/", handlers.Login)
	user.GET("/register/", handlers.Register)
	user.GET("/", handlers.GetUserById)

}

// 初始化 Hertz API 及 Router
func main() {

	h := InitHertz()
	// h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
	// ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	// })

	registerGroup(h)

	h.Spin()
}
