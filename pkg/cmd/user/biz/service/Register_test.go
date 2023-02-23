package service

import (
	"context"
	user "github.com/AgSword/simpleDouyin/kitex_gen/user"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value
	req := &user.UserRegisterRequest{Username: "liyinjian", Password: "liyinjian"}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	print(resp)
	// todo: edit your unit test

}
