package main

import (
	"context"
	publish "github.com/AgSword/simpleDouyin/kitex_gen/publish"
	"github.com/AgSword/simpleDouyin/cmd/publish/service"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	//claim, err := Jwt.ParseToken(req.Token)
	//uid := int(claim.Id)
	uid := 1
	if err := service.PublishAction(req, uid); err != nil {
		return &publish.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  nil,
		}, err
	}

	return &publish.PublishActionResponse{
		StatusCode: 0,
		StatusMsg:  nil,
	}, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	list, err := service.PublishList(req)
	if err != nil {
		return &publish.PublishListResponse{
			StatusCode: -1,
			StatusMsg:  nil,
			VideoList:  nil,
		}, err
	}

	return &publish.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		VideoList:  list,
	}, err
}
