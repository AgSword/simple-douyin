package rpc

import (
	"context"
	// "fmt"
	// "time"
	"errors"

	// favorite "github.com/Karry-Almond/favorite/kitex_gen/api"
	// favoritesrv "github.com/Karry-Almond/favorite/kitex_gen/api/favorite"
	// "github.com/Karry-Almond/favorite/kitex_gen/api"
	// "github.com/Karry-Almond/favorite/kitex_gen/api/favorite"
	"github.com/AgSword/simpleDouyin/kitex_gen/favorite"
	"github.com/AgSword/simpleDouyin/kitex_gen/favorite/favoriteservice"
	"github.com/cloudwego/kitex/client"
)

var favoriteClient favoriteservice.Client

// Favorite RPC 客户端初始化
func init() {

	c, err := favoriteservice.NewClient("Favorite", client.WithHostPorts("127.0.0.1:9999"))
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// 传递 点赞操作 的上下文, 并获取 RPC Server 端的响应.
func FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp, err = favoriteClient.Action(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New("Action Failed")
	}
	return resp, nil
}

// 传递 获取点赞列表操作 的上下文, 并获取 RPC Server 端的响应.
// func FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
// 	resp, err = favoriteClient.FavoriteList(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != 0 {
// 		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
// 	}
// 	return resp, nil
// }
