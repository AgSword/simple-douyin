package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"

	"github.com/AgSword/simple-douyin/cmd/user/biz/dal/mysql"
	user "github.com/AgSword/simple-douyin/kitex_gen/user"
)

type GetUserByIdService struct {
	ctx context.Context
} // NewGetUserByIdService new GetUserByIdService
func NewGetUserByIdService(ctx context.Context) *GetUserByIdService {
	return &GetUserByIdService{ctx: ctx}
}

// Run create note info
func (s *GetUserByIdService) Run(req *user.UserRequest) (resp *user.UserResponse, err error) {
	// Finish your business logic.
	// 从数据库中根据id获取user数据，封装到resp中
	ids := make([]int32, 0)
	ids[0] = int32(req.UserId)
	users, err := mysql.GetUserByIds(s.ctx, ids)
	if err != nil || len(users) != 1 {
		return nil, err
	}

	msg := "ok"
	userVo := user.User{Id: users[0].ID, Name: users[0].UserName, FollowCount: users[0].FollowerCount, FollowerCount: users[0].FollowerCount}
	return user.UserResponse{StatusCode: int32(codes.OK), StatusMsg: &msg, User: users[0]}, nil
	return
}
