package rpc

import (
	"context"
	"errors"
	// "log"
	// "time"

	"github.com/AgSword/simpleDouyin/kitex_gen/relation"
	relation2 "github.com/AgSword/simpleDouyin/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/client"
	//"github.com/cloudwego/kitex/client/callopt"
	//"github.com/AgSword/simpleDouyin/cmd/api/handlers"
	
)

var relationClient relation2.Client

func init() {

	c, err := relation2.NewClient("Relation A", client.WithHostPorts("127.0.0.1:5555"))
	if err != nil {
		panic(err)
	}
	relationClient = c
}
//RelationActionParam
func RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp, err = relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil, errors.New("RelationAction Failed")
	}
	return resp, nil
}

func RelationFollowList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp, err = relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil, errors.New("RelationFollowList Failed")
	}
	return resp, nil
}

func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp, err = relationClient.RelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil, errors.New("RelationFollowerList Failed")
	}
	return resp, nil
}


