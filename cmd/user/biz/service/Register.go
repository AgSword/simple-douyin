package service

import (
	"context"
	"github.com/AgSword/simpleDouyin/cmd/user/biz/dal/mysql"
	"github.com/AgSword/simpleDouyin/pkg/md5"

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
	userIDb := mysql.User{UserName: req.Username, FollowingCount: 0, FollowerCount: 0}
	err = mysql.CreateUser(s.ctx, &userIDb)
	if err != nil {
		return nil, err
	}
	// 构造密码
	pw := md5.MD5(string(userIDb.ID) + req.Password + string(userIDb.ID))
	userIDb.Password = pw
	// 将密码保存到数据库
	err = mysql.UpdateUserPassword(s.ctx, &userIDb)
	if err != nil {
		return nil, err
	}
	loginRequest := user.UserLoginRequest{Username: userIDb.UserName, Password: userIDb.Password}
	loginResponse, err := NewLoginService(s.ctx).Run(&loginRequest)
	if err != nil {
		return nil, err
	}
	resp.UserId = loginResponse.UserId
	resp.StatusCode = loginResponse.StatusCode
	resp.StatusMsg = loginResponse.StatusMsg
	resp.Token = loginResponse.Token
	return resp, err
}
