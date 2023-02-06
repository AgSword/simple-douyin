package service

import (
	"context"
	"errors"
	"github.com/AgSword/simpleDouyin/cmd/user/biz/dal/mysql"
	"github.com/AgSword/simpleDouyin/pkg/md5"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"

	user "github.com/AgSword/simpleDouyin/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// Finish your business logic.
	// 检查是否存在这个用户名
	username := req.Username
	users, err := mysql.GetUserByUsername(s.ctx, username)
	if err != nil {
		return nil, err
	}
	// 用户不存在
	if len(users) == 0 {
		return nil, errors.New("the user no exist in db")
	}
	if len(users) != 1 {
		return nil, errors.New("have mul users whose name are the same in db")
	}
	// 判断密码是否正确
	md5_src := string(users[0].ID) + req.Password + string(users[0].ID)
	md5_dest := md5.MD5(md5_src)
	if md5_dest != users[0].Password {
		return nil, errors.New("username or password error")
	}
	// 因为目录安排的原因，这个函数只能到判断用户名密码是否正确这一步，在handler层返回token
	resp.UserId = users[0].ID
	resp.StatusCode = int32(codes.OK)
	msg := "ok"
	resp.StatusMsg = &msg
	return resp, nil

}
