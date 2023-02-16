package main

//import (
//	"context"
//	"favorite/kitex_gen/api"
//	"favorite/kitex_gen/api/favorite"
//	"log"
//	"time"
//
//	"github.com/cloudwego/kitex/client"
//	"github.com/cloudwego/kitex/client/callopt"
//)
//
//func main() {
//	c, _ := favorite.NewClient("Favorite", client.WithHostPorts("0.0.0.0:8888"))
//	req := &api.DouyinFavoriteListRequest{UserId: 1}
//	resp, err := c.List(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
//	if err != nil {
//		log.Println("ERROR")
//		log.Println(err)
//	}
//	log.Println("RESP")
//	log.Println(resp)
//
//}
