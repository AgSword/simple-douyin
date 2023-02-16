package handlers

import (
	"context"
	"strconv"

	"github.com/AgSword/simpleDouyin/cmd/api/rpc"
	// "github.com/AgSword/simple-douyin/kitex_gen/user"
	"github.com/AgSword/simpleDouyin/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

func Login(ctx context.Context, c *app.RequestContext) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		msg := "Invalid UserName or PassWord"
		SendResponse(c, &user.UserLoginResponse{
			StatusCode: int32(1),
			StatusMsg:  &msg,
			UserId:     0,
			Token:      "",
		})
		return
	}

	resp, err := rpc.Login(ctx, &user.UserLoginRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		msg := "Invalid UserName or PassWord"
		SendResponse(c, &user.UserLoginResponse{
			StatusCode: int32(2),
			StatusMsg:  &msg,
			UserId:     0,
			Token:      "",
		})
		return
	}
	SendResponse(c, resp)
}

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		msg := "Invalid UserName or PassWord"
		SendResponse(c, &user.UserRegisterResponse{
			StatusCode: int32(1),
			StatusMsg:  &msg,
			// UserId:		0,
			// Token:		"",
		})
		return
	}
	resp, err := rpc.Register(ctx, &user.UserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		msg := "Invalid UserName or PassWord"
		SendResponse(c, &user.UserRegisterResponse{
			StatusCode: int32(2),
			StatusMsg:  &msg,
			// UserId:		0,
			// Token:		"",
		})
		return
	}
	SendResponse(c, resp)

}

func GetUserById(ctx context.Context, c *app.RequestContext) {
	var userVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		msg := "Invalid UserID"
		SendResponse(c, &user.UserResponse{
			StatusCode: int32(1),
			StatusMsg:  &msg,
			// UserId:		0,
			// Token:		"",
		})
		return
	}
	userVar.UserId = int64(userid)
	userVar.Token = c.Query("token")

	if len(userVar.Token) == 0 || userVar.UserId < 0 {
		msg := "Invalid Token"
		SendResponse(c, &user.UserResponse{
			StatusCode: int32(2),
			StatusMsg:  &msg,
			// UserId:		0,
			// Token:		"",
		})
		return
	}

	resp, err := rpc.GetUserById(ctx, &user.UserRequest{
		UserId: userVar.UserId,
		Token:  userVar.Token,
	})
	if err != nil {
		msg := "Invalid Token"
		SendResponse(c, &user.UserResponse{
			StatusCode: int32(3),
			StatusMsg:  &msg,
			// UserId:		0,
			// Token:		"",
		})
		return
	}
	SendResponse(c, resp)
}
