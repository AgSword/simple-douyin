package service

import (
	"context"
	"errors"
	"github.com/AgSword/simpleDouyin/dal/mysql"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"

	user "github.com/AgSword/simpleDouyin/kitex_gen/user"
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

	users, err := mysql.GetUserByIds(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, errors.New("user no exist")
	}
	msg := "ok"
	userVo := user.User{Id: users.ID, Name: users.Name, FollowCount: &users.FollowerCount, FollowerCount: &users.FollowerCount}
	return &user.UserResponse{StatusCode: int32(codes.OK), StatusMsg: &msg, User: &userVo}, nil
}
