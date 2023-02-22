package handlers

import (
	"bytes"
	"context"
	"io"
	//"time"

	"github.com/AgSword/simpleDouyin/cmd/api/rpc"
	"github.com/AgSword/simpleDouyin/kitex_gen/publish"

	//"log"
	"strconv"

	//"errs"

	"github.com/cloudwego/hertz/pkg/app"
)

// 传递 发布视频操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var paramVar PublishActionparam
	token := c.PostForm("token")
	title := c.PostForm("title")

	fileHeader, err := c.Request.FormFile("data")

    if err != nil {
		msg := "传参解析错误"
		SendResponse(c, &publish.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
			//关注一下结果
		})
		return
	}

	file, err := fileHeader.Open()
    if err != nil {
		msg := "打开文件失败"
		SendResponse(c, &publish.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
			//关注一下结果
		})
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		msg := "文件转换失败"
		SendResponse(c, &publish.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
			//关注一下结果
		})
		return
	}

	paramVar.Token = token
	paramVar.Title = title

	resp, err := rpc.PublishAction(ctx, &publish.PublishActionRequest{
		Title: paramVar.Title,
		Token: paramVar.Token,
		Data:  buf.Bytes(),
	})
	if err != nil {
		msg := "PublishAction rpc这步错了"
		SendResponse(c, &publish.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
		})
		return
	}
	SendResponse(c, resp)
}

// 传递 获取视频列表操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func PublishList(ctx context.Context, c *app.RequestContext) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		msg := "publishlist rpc这步错了"
		SendResponse(c, &publish.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
		})
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		msg := "userid 或 token这步错了"
		SendResponse(c, &publish.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
		})
		return
	}

	resp, err := rpc.PublishList(ctx, &publish.PublishListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		msg := "publishList rpc这步错了"
		SendResponse(c, &publish.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
		})
		return
	}
	SendResponse(c, resp)
}

