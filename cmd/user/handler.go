package main

import (
	"context"
	service2 "github.com/AgSword/simpleDouyin/cmd/user/service"
	"github.com/AgSword/simpleDouyin/pkg/jwt"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"

	user "github.com/AgSword/simpleDouyin/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
// user服务api的实现类
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {

	if len(req.Username) == 0 || len(req.Password) == 0 {
		msg := "invalid argument"
		return &user.UserRegisterResponse{StatusCode: int32(codes.InvalidArgument), StatusMsg: &msg}, nil
	}
	resp, err = service2.NewRegisterService(ctx).Run(req)
	// register service error
	if err != nil {
		msg := "register service error"
		return &user.UserRegisterResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	token, err := Jwt.CreateToken(jwt.CustomClaims{ID: resp.UserId})

	if err != nil {
		msg := "create token fail"
		return &user.UserRegisterResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	resp.Token = token
	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {

	if len(req.Username) == 0 || len(req.Password) == 0 {
		msg := "invalid argument"
		return &user.UserLoginResponse{StatusCode: int32(codes.InvalidArgument), StatusMsg: &msg}, nil
	}
	resp, err = service2.NewLoginService(ctx).Run(req)
	if err != nil {
		msg := "login service error"
		return &user.UserLoginResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	token, err := Jwt.CreateToken(jwt.CustomClaims{ID: resp.UserId})

	if err != nil {
		msg := "create token fail"
		return &user.UserLoginResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	resp.Token = token
	return resp, err
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {

	if len(req.Token) == 0 {
		msg := "invalid argument"
		return &user.UserResponse{StatusCode: int32(codes.InvalidArgument), StatusMsg: &msg}, nil
	}
	resp, err = service2.NewGetUserByIdService(ctx).Run(req)
	if err != nil {
		msg := "getUserById service error"
		return &user.UserResponse{StatusCode: int32(codes.Unimplemented), StatusMsg: &msg}, nil
	}
	return resp, err
}
