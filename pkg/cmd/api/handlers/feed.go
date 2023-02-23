package handlers

import (
	"context"
	"time"

	"github.com/AgSword/simpleDouyin/cmd/api/rpc"
	"github.com/AgSword/simpleDouyin/kitex_gen/feed"

	"log"
	"strconv"
	//"errs"

	"github.com/cloudwego/hertz/pkg/app"
)

// 视频流接口  /douyin/feed/

func GetUserFeed(ctx context.Context, c *app.RequestContext)  {
	// TODO: Your code here...
	var feedVar Feedparam
	var latestTime int64
	var err error
	latestTimeStr := c.Query("latest_time")
	if len(latestTimeStr)==0 || latestTimeStr=="0"{
		latestTimeStr=time.Now().Format("2006-01-02 15:04:05")
		latestTimeint:=latestTimeStr[:4]+latestTimeStr[5:7]+
		latestTimeStr[8:10]+latestTimeStr[11:13]+
		latestTimeStr[14:16]+latestTimeStr[17:19]
		latestTime,err=strconv.ParseInt(latestTimeint, 10, 64)
		if err!=nil{
			log.Println(err.Error())
			log.Println("latestTime错误")
		}
	}else{
		latestTime,err=strconv.ParseInt(latestTimeStr, 10, 64)
	}

	feedVar.Token = c.Query("token")
	feedVar.Latest_time=latestTime

    // if latest_time为空  置为当前时间
	resp, err := rpc.GetUserFeed(ctx, &feed.FeedRequest{
		LatestTime: &feedVar.Latest_time,
		Token:  &feedVar.Token,	
	})
	if err != nil {
		msg := "Invalid Token feed"
		SendResponse(c, &feed.FeedResponse{
			StatusCode: int32(3), 
			StatusMsg:  &msg,	

		})
		return
	}
	SendResponse(c, resp)

}

