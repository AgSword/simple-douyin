package service

import (
	"context"
	"github.com/AgSword/simpleDouyin/dal/mysql"
	"github.com/AgSword/simpleDouyin/pkg/md5"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"strconv"

	user "github.com/AgSword/simpleDouyin/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// Finish your business logic.
	// 在数据库中创建用户信息,先添加用户名添加记录行，后添加密码
	userInDb := mysql.User{Name: req.Username, FollowCount: 0, FollowerCount: 0}
	err = mysql.CreateUser(s.ctx, &userInDb)
	if err != nil {
		return nil, err
	}
	// 构造密码
	pw := md5.MD5(strconv.FormatInt(userInDb.ID, 10) + req.Password + strconv.FormatInt(userInDb.ID, 10))
	userInDb.Password = pw
	// 将密码保存到数据库
	err = mysql.UpdateUserPassword(s.ctx, &userInDb)
	if err != nil {
		return nil, err
	}
	msg := "ok"
	return &user.UserRegisterResponse{UserId: userInDb.ID, StatusMsg: &msg, StatusCode: int32(codes.OK)}, err
}
