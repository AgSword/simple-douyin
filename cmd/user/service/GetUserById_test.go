package service

import (
	"context"
	"github.com/AgSword/simpleDouyin/dal/mysql"
	user "github.com/AgSword/simpleDouyin/kitex_gen/user"
	"github.com/AgSword/simpleDouyin/pkg/jwt"
	"testing"
)

func TestGetUserById_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewGetUserByIdService(ctx)
	// init req and assert value
	newJWT := jwt.NewJWT([]byte("signkey"))
	var id int64 = 8
	token, err2 := newJWT.CreateToken(jwt.CustomClaims{ID: id})
	if err2 != nil {
		panic(err2)
	}
	req := &user.UserRequest{UserId: id, Token: token}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
