package main

import (
	// "context"
	// "errors"
	// "fmt"

	"github.com/AgSword/simpleDouyin/cmd/api/handlers"
	//"google.golang.org/genproto/googleapis/streetview/publish/v1"
	// "github.com/Liujony/tiktokapi/cmd/api/rpc"
	"github.com/cloudwego/hertz/pkg/common/config"
	//"github.com/cloudwego/hertz/pkg/network/netpoll"
	// "github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/app/server/registry"
	//"github.com/hertz-contrib/registry/etcd"
)

var (
	ServiceAddr = "0.0.0.0:8080"
) 

// 初始化 Hertz
func InitHertz() *server.Hertz {
	opts := []config.Option{server.WithHostPorts(ServiceAddr)}

	opts =append(opts,server.WithMaxRequestBodySize(104857600))
	h := server.Default(opts...)
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

	publish:=douyin.Group("/publish")
	publish.POST("/action/",handlers.PublishAction)
	publish.GET("/list",handlers.PublishList)

	relation:=douyin.Group("/relation")
	relation.POST("/action/",handlers.RelationAction)
	relation.GET("/follow/list",handlers.RelationFollowList)
	relation.GET("/follower/list",handlers.RelationFollowerList)

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
