package service

import (
	"context"
	"github.com/AgSword/simpleDouyin/dal/mysql"
	user "github.com/AgSword/simpleDouyin/kitex_gen/user"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.UserLoginRequest{Username: "liyinjian", Password: "liyinjian"}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test
	println(resp.UserId)
	println(resp.Token)
}
