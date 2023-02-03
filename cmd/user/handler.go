package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"

	"github.com/AgSword/simple-douyin/cmd/user/biz/service"
	user "github.com/AgSword/simple-douyin/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
// user服务api的实现类
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// 参数校验，需要调整 TODO
	if len(req.Username) == 0 || len(req.Password) == 0 {
		msg := "invalid argument"
		return &user.UserRegisterResponse{StatusCode: int32(codes.InvalidArgument), StatusMsg: &msg}, nil
	}
	resp, err = service.NewRegisterService(ctx).Run(req)
	// register service error
	if err != nil {
		msg := "register service error"
		return &user.UserRegisterResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// 参数校验，需要调整 TODO
	if len(req.Username) == 0 || len(req.Password) == 0 {
		msg := "invalid argument"
		return &user.UserLoginResponse{StatusCode: int32(codes.InvalidArgument), StatusMsg: &msg}, nil
	}
	resp, err = service.NewLoginService(ctx).Run(req)
	if err != nil {
		msg := "login service error"
		return &user.UserLoginResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	return resp, err
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	// 参数校验，需要调整 TODO
	if len(req.Token) == 0 {
		msg := "invalid argument"
		return &user.UserResponse{StatusCode: int32(codes.InvalidArgument), StatusMsg: &msg}, nil
	}
	resp, err = service.NewGetUserByIdService(ctx).Run(req)
	if err != nil {
		msg := "getUserById service error"
		return &user.UserResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	return resp, err
}
