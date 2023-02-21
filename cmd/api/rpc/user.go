package rpc

import (
	"context"
	"fmt"
	// "fmt"
	// "time"
	"errors"

	// "github.com/AgSword/simple-douyin/kitex_gen/user"
	// "github.com/AgSword/simple-douyin/kitex_gen/user/userservice"
	"github.com/AgSword/simpleDouyin/kitex_gen/user"
	"github.com/AgSword/simpleDouyin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

// User RPC 客户端初始化
// func initUserRpc(Config *ttviper.Config) {
func init() {

	c, err := userservice.NewClient("User", client.WithHostPorts("127.0.0.1:7777"))
	if err != nil {
		panic(err)
	}
	userClient = c
}

// 传递 注册操作 的上下文, 并获取 RPC Server 端的响应.
func Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp, err = userClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil, errors.New("Action Failed")
	}
	return resp, nil
}

// 传递 登录操作 的上下文, 并获取 RPC Server 端的响应.
func Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp, err = userClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil, errors.New("Action Failed")
	}
	fmt.Println(err)
	return resp, nil
}

// 传递 获取用户信息操作 的上下文, 并获取 RPC Server 端的响应.
func GetUserById(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp, err = userClient.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		fmt.Println(resp.StatusCode)
		return nil, errors.New("Action Failed")
	}
	return resp, nil
}
