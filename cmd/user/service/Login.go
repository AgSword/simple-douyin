package service

import (
	"context"
	"errors"
	"github.com/AgSword/simpleDouyin/dal/mysql"
	"github.com/AgSword/simpleDouyin/pkg/md5"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"strconv"

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
	userInDb, err := mysql.GetUserByUsername(s.ctx, username)
	if err != nil {
		return nil, err
	}
	// 用户不存在
	if userInDb == nil {
		return nil, errors.New("the user no exist in db")
	}
	// 判断密码是否正确
	md5_src := strconv.FormatInt(userInDb.ID, 10) + req.Password + strconv.FormatInt(userInDb.ID, 10)
	md5_dest := md5.MD5(md5_src)
	if md5_dest != userInDb.Password {
		return nil, errors.New("username or password error")
	}

	// 因为目录安排的原因，这个函数只能到判断用户名密码是否正确这一步，在handler层返回token
	msg := "ok"

	return &user.UserLoginResponse{UserId: userInDb.ID, StatusMsg: &msg, StatusCode: int32(codes.OK)}, nil

}
