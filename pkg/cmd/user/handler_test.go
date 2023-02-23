package main

import (
	"context"
	"github.com/AgSword/simpleDouyin/kitex_gen/user"
	"github.com/AgSword/simpleDouyin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {

	newClient, err := userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	request := user.UserRegisterRequest{Username: "ggg", Password: "ggg"}
	response, err := newClient.Register(context.Background(), &request, callopt.WithRPCTimeout(3*time.Second))
	println(response.UserId, response.StatusCode, response.StatusMsg, response.Token)
}

func TestLogin(t *testing.T) {
	newClient, err := userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	request := user.UserLoginRequest{Username: "ggg", Password: "ggg"}
	response, err := newClient.Login(context.Background(), &request, callopt.WithRPCTimeout(3*time.Second))
	println(response.UserId, response.StatusCode, response.StatusMsg, response.Token)
}

func TestUserServiceImpl_GetUserById(t *testing.T) {
	newClient, err := userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	request := user.UserRequest{UserId: 21, Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MjF9.8QPUAtQ7ZWlWWFsxFWEVyYs_3jXjfiUuRIlr0mENpJg"}
	response, err := newClient.GetUserById(context.Background(), &request, callopt.WithRPCTimeout(3*time.Second))
	println(response.User.Name, response.User.Id, response.User.FollowCount, response.User.FollowerCount)
	println(response.StatusCode, response.StatusMsg)
}
