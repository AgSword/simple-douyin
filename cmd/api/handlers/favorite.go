package handlers

import (
	"context"
	"github.com/AgSword/simpleDouyin/cmd/api/rpc"
	"github.com/AgSword/simpleDouyin/kitex_gen/favorite"
	"log"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// 传递 点赞操作 的上下文至 Favorite 服务的 RPC 客户端, 并获取相应的响应.
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var paramVar FavoriteActionParam
	// token := c.Query("token")
	user_id := c.Query("user_id")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	// 如果需要在这里验证JWT
	// 添加userid

	uid, err := strconv.Atoi(user_id)
	if err != nil {
		log.Println(err.Error())
		SendResponse(c, &favorite.DouyinFavoriteActionResponse{StatusCode: int32(1), StatusMsg: "Invalid UserId"})
		return
	}

	vid, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, &favorite.DouyinFavoriteActionResponse{StatusCode: int32(2), StatusMsg: "Invalid VideoId"})
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		// SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		SendResponse(c, &favorite.DouyinFavoriteActionResponse{StatusCode: int32(3), StatusMsg: "Invalid ActionType"})
		return
	}

	// paramVar.Token = token
	paramVar.UserId = int64(uid)
	paramVar.VideoId = int64(vid)
	paramVar.ActionType = int32(act)

	resp, err := rpc.FavoriteAction(ctx, &favorite.DouyinFavoriteActionRequest{
		VideoId: paramVar.VideoId,
		// Token:      paramVar.Token,
		UserId:     paramVar.UserId,
		ActionType: paramVar.ActionType,
	})
	if err != nil {
		// SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		SendResponse(c, &favorite.DouyinFavoriteActionResponse{StatusCode: int32(4), StatusMsg: "RPC Resp Failed"})
		return
	}
	SendResponse(c, resp)
}

// 传递 获取点赞列表操作 的上下文至 Favorite 服务的 RPC 客户端, 并获取相应的响应.
// func FavoriteList(ctx context.Context, c *app.RequestContext) {
// 	var paramVar UserParam
// 	userid, err := strconv.Atoi(c.Query("user_id"))
// 	if err != nil {
// 		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
// 		return
// 	}
// 	paramVar.UserId = int64(userid)
// 	paramVar.Token = c.Query("token")

// 	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
// 		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
// 		return
// 	}

// 	resp, err := rpc.FavoriteList(ctx, &favorite.DouyinFavoriteListRequest{
// 		UserId: paramVar.UserId,
// 		Token:  paramVar.Token,
// 	})
// 	if err != nil {
// 		SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
// 		return
// 	}
// 	SendResponse(c, resp)
// }
