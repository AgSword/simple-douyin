package main

import (
	favorite "github.com/AgSword/simpleDouyin/kitex_gen/favorite/favoriteservice"
	"log"
	"net"

	"github.com/AgSword/simpleDouyin/cmd/favorite/packages/favoriteDB"
	"github.com/cloudwego/kitex/server"
)

func main() {
	favoriteDB.Init()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	svr := favorite.NewServer(new(FavoriteImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
