package main

import (
	"log"
	relation "github.com/AgSword/simpleDouyin/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/AgSword/simpleDouyin/cmd/relation/packages/relationDB"
	"net"
	"github.com/AgSword/simpleDouyin/pkg/jwt"
	//github.com/AgSword/simpleDouyin
)

var Jwt *jwt.JWT

func main() {
	relationDB.Init()
	Jwt = jwt.NewJWT([]byte("signkey"))
	var opts []server.Option
	addr, err := net.ResolveTCPAddr("tcp", ":5555")
	if err != nil {
		panic(err)
	}
    opts=append(opts,server.WithServiceAddr(addr))
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "publish",
	}))
	svr := relation.NewServer(new(RelationServiceImpl),opts...)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
