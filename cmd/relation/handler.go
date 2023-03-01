package main

import (
	"context"
	relation "github.com/AgSword/simpleDouyin/kitex_gen/relation"
	"github.com/AgSword/simpleDouyin/cmd/relation/packages/relationDB"
	"fmt"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// 类型是关注请求


	token,err:=Jwt.ParseToken(req.Token)
	if err!=nil{
		fmt.Println(err)
	}
	// fmt.Println(token.ID)
	if req.ActionType == 1 {
		if _, err = relationDB.NewRelation(token.ID, req.ToUserId); err != nil {
			return &relation.RelationActionResponse{
				// status_code = 1 表示关注失败
				StatusCode: 1,
					StatusMsg:  err.Error()}, nil
		}
	} else if req.ActionType == 2 {
		if _, err = relationDB.CancelRelation(req.UserId, req.ToUserId); err != nil {
			return &relation.RelationActionResponse{
				// status_code = 2 表示取消关注失败
				StatusCode: 2,
				StatusMsg:  err.Error()}, nil
		}
	} else {
		return &relation.RelationActionResponse{
			// status_code = 3 表示类型信息错误
			StatusCode: 3,
			StatusMsg:  "action_type error"}, nil
	}

	return &relation.RelationActionResponse{
		// status_code = 0 表示操作成功
		StatusCode: 0,
		StatusMsg:  "action success"}, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowListResponse, err error) {
	_, list, err := relationDB.GetFollowList(req.UserId)
	if err != nil {
		return &relation.RelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
			UserList:   list,
		}, nil
	}
	return &relation.RelationFollowListResponse{
		// status_code = 0 表示操作成功
		StatusCode: 0,
		StatusMsg:  "RelationFollowList success",
		UserList:   list,
		}, nil
		
}
	

	


// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	_, list, err := relationDB.GetFollowerList(req.UserId)
	if err != nil {
		str := err.Error()
		return &relation.RelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  str,
			UserList:   nil,
		}, nil
	}
	return &relation.RelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   list,
	}, nil

}
