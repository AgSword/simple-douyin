package main

import (
	// "context"
	// "errors"
	// "fmt"

	"github.com/AgSword/simpleDouyin/cmd/api/handlers"
	// "github.com/Liujony/tiktokapi/cmd/api/rpc"
	"github.com/cloudwego/hertz/pkg/common/config"
	// "github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/registry/etcd"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
)

var (
	ServiceAddr = "0.0.0.0:8080"
) 

// 初始化 Hertz
func InitHertz() *server.Hertz {

	opts := []config.Option{server.WithHostPorts(ServiceAddr)}
	r, _ := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
		// if err != nil {
		// 	hlog.Fatal(err)
		// }
		opts = append(opts, server.WithRegistry(r, &registry.Info{
			//ServiceName: ServiceName,
			//Addr:        utils.NewNetAddr("tcp", ServiceAddr),
			//Weight:      10,
			//Tags:        nil,
		}))


	h := server.Default(server.WithHostPorts(ServiceAddr))
	return h

	
}

// 注册 Router组
func registerGroup(h *server.Hertz) {
	douyin := h.Group("/douyin")
	//douyin.GET("/",handlers.GetUserFeed)

	favorite := douyin.Group("/favorite")
	favorite.POST("/action/", handlers.FavoriteAction)

	user := douyin.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	user.GET("/", handlers.GetUserById)

	feed:=douyin.Group("/feed")
	feed.GET("/", handlers.GetUserFeed)

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
