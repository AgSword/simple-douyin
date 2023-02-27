package handlers

import (
	// 	"bytes"
	"context"
	"strconv"
	// 	"io"
	// 	//"time"
	//"github.com/AgSword/simpleDouyin/cmd/relation"
	"github.com/AgSword/simpleDouyin/cmd/api/rpc"
	"github.com/AgSword/simpleDouyin/kitex_gen/relation"

	// 	"github.com/AgSword/simpleDouyin/kitex_gen/relation"

	// 	//"log"
	// 	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func RelationAction(ctx context.Context, c *app.RequestContext) {
	var paramVar RelationActionParam
	paramVar.Token= c.Query("token")
	paramVar.Action_type,_=strconv.Atoi(c.Query("action_type"))
	paramVar.To_user_id,_=strconv.Atoi(c.Query("to_user_id"))

    
	resp, err := rpc.RelationAction(ctx, &relation.RelationActionRequest{
		Token: paramVar.Token,
		ActionType: int32(paramVar.Action_type),
		ToUserId:  int64(paramVar.To_user_id),

	})

    if err != nil {
		msg := "RelationAction错误"
		SendResponse(c, &relation.RelationActionResponse{
			StatusCode: -1,
			StatusMsg:  msg,
			//关注一下结果
		})
		return
	}
 
	SendResponse(c,resp)

}

func RelationFollowList(ctx context.Context, c *app.RequestContext) {

	var paramVar RelationFollowListParam
	paramVar.Token= c.Query("token")  //和c.query区别
	paramVar.User_id,_=strconv.Atoi(c.Query("user_id"))

	resp, _ := rpc.RelationFollowList(ctx, &relation.RelationFollowerListRequest{
		Token: paramVar.Token,
		UserId:  int64(paramVar.User_id),
	})

	SendResponse(c,resp)

}

func RelationFollowerList(ctx context.Context,c *app.RequestContext){
	var paramVar RelationFollowListParam
	paramVar.Token= c.Query("token")
	paramVar.User_id,_=strconv.Atoi(c.Query("user_id"))

	resp, _ := rpc.RelationFollowerList(ctx, &relation.RelationFollowerListRequest{
		Token: paramVar.Token,
		UserId:  int64(paramVar.User_id),
	})

	SendResponse(c,resp)

}