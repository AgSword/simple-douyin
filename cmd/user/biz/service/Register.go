package service

import (
	"context"

	user "github.com/AgSword/simple-douyin/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// Finish your business logic.

	return
}
