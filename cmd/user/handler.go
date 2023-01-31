package main

import (
	"context"

	"github.com/AgSword/simple-douyin/cmd/user/biz/service"
	user "github.com/AgSword/simple-douyin/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
// user服务api的实现类
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = new(user.UserRegisterResponse)
	if len(req.Username) == 0 || len(req.Password) == 0 {

	}
	// 参数校验
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp, err = service.NewGetUserByIdService(ctx).Run(req)

	return resp, err
}
