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
	ids := make([]int64, 1)
	ids[0] = req.UserId
	users, err := mysql.GetUserByIds(s.ctx, ids)
	if err != nil {
		return nil, err
	}
	if len(users) != 1 {
		return nil, errors.New("user no exist or have mul users")
	}
	msg := "ok"
	userVo := user.User{Id: users[0].ID, Name: users[0].Name, FollowCount: &users[0].FollowerCount, FollowerCount: &users[0].FollowerCount}
	return &user.UserResponse{StatusCode: int32(codes.OK), StatusMsg: &msg, User: &userVo}, nil
}
