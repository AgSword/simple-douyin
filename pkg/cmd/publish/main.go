package main

import (
	"log"
	publish "github.com/AgSword/simpleDouyin/kitex_gen/publish/publishservice"
	"github.com/AgSword/simpleDouyin/cmd/publish/model/db"

	//"github.com/cloudwego/kitex/internal/server"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"net"
)

func main() {
	db.Init()
	var opts []server.Option
	addr, err := net.ResolveTCPAddr("tcp", ":6666")
	if err != nil {
		panic(err)
	}
    opts=append(opts,server.WithServiceAddr(addr))
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "publish",
	}))
	svr := publish.NewServer(new(PublishServiceImpl),opts...)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
