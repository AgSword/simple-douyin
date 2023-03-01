package main

import (
	"context"

	//"github.com/AgSword/simpleDouyin/pkg/jwt"
	//service2 "github.com/AgSword/simpleDouyin/cmd/feed/service"
	"github.com/AgSword/simpleDouyin/cmd/feed/service"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
    "fmt"
	feed "github.com/AgSword/simpleDouyin/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type  FeedServiceImpl struct{}

// 视频流接口  /douyin/feed/
//得改
func (s *FeedServiceImpl) GetUserFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	if req.Token == nil {
		msg := "invalid argument"
		return &feed.FeedResponse{StatusCode: int32(codes.InvalidArgument), StatusMsg: &msg}, nil
	}
	fmt.Println(Jwt)
	token,err:=Jwt.ParseToken(*req.Token)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(token.ID)
	resp, err = service.NewGetUserFeedService(ctx).Run(req)

	if err != nil {
		msg := "getUserById service error"
		return &feed.FeedResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	return resp, err

}
