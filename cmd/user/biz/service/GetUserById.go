package service

import (
	"context"

	user "github.com/AgSword/simple-douyin/kitex_gen/user"
)

type GetUserByIdService struct {
	ctx context.Context
} // NewGetUserByIdService new GetUserByIdService
func NewGetUserByIdService(ctx context.Context) *GetUserByIdService {
	return &GetUserByIdService{ctx: ctx}
}

// Run create note info
func (s *GetUserByIdService) Run(req *user.UserRequest) (resp *user.UserResponse, err error) {
	// Finish your business logic.

	return
}
