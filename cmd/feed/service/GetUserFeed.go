package service

import (
	"context"
	"strconv"
	"time"

	//"errors"
	"fmt"

	"github.com/AgSword/simpleDouyin/dal/mysql"
	feed "github.com/AgSword/simpleDouyin/kitex_gen/feed"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	//mysql2 "github.com/AgSword/simpleDouyin/dal/mysql"
	//user "github.com/AgSword/simpleDouyin/kitex_gen/user"
	//"errors"

)

type GetUserFeedService struct {
	ctx context.Context
}

func NewGetUserFeedService(ctx context.Context) *GetUserFeedService {
	return &GetUserFeedService{ctx: ctx}
}

func (s *GetUserFeedService) Run(req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {

	// 从数据库中根据last_time获取video数据，封装到resp中
	latestTime := req.LatestTime
	videolist, err := mysql.GetUserFeeds(s.ctx, latestTime)
	if err != nil {
		return nil, err
	}
	if len(videolist) == 0 {
		fmt.Println("当前无视频！")
	}

	videolistVo := make([]*feed.Video, len(videolist))
	ids := make([]int64, len(videolist))

	for i := 0; i < len(videolist); i++ {
		ids[i] = videolist[i].UserId
		users, err := mysql.GetUserByIds(s.ctx, ids[i])
		if err != nil {
			return nil, err
		}
		// if len(users) != 1 {
		// 	return nil, errors.New("user no exist or have mul users")
		// }

		userVo := feed.User{Id: users.ID, Name: users.Name, FollowCount: &users.FollowerCount, FollowerCount: &users.FollowerCount}
		videolistVo[i] = &feed.Video{Id: videolist[i].ID, Author: &userVo,
			PlayUrl: videolist[i].PlayUrl, CoverUrl: videolist[i].CoverUrl,
			FavoriteCount: videolist[i].FavoriteCount, CommentCount: videolist[i].CommentCount,
			IsFavorite: false, Title: videolist[i].Title}
	}
	msg := "ok"

	nextTime := videolist[len(videolist)-1].CreatedAt.Unix()
	nextTimestrs := time.Unix(nextTime, 0).Format("2006-01-02 15:04:05")
	nextTimeint:=nextTimestrs[:4]+nextTimestrs[5:7]+
	nextTimestrs[8:10]+nextTimestrs[11:13]+
	nextTimestrs[14:16]+nextTimestrs[17:19]
	nextTime,_=strconv.ParseInt(nextTimeint, 10, 64)

	return &feed.FeedResponse{StatusCode: int32(codes.OK),
		 StatusMsg: &msg, VideoList: videolistVo, 
		NextTime: &nextTime}, nil

}
