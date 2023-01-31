package service

import (
	"context"
	user "github.com/AgSword/simple-douyin/kitex_gen/user"
	"testing"
)

func TestGetUserById_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetUserByIdService(ctx)
	// init req and assert value

	req := &user.UserRequest{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
