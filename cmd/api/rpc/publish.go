package rpc

import(
	"github.com/AgSword/simpleDouyin/kitex_gen/publish/publishservice"
	"github.com/cloudwego/kitex/client"
	"github.com/AgSword/simpleDouyin/kitex_gen/publish"
	"context"
	"errors"
)

var publishClient publishservice.Client


func init() {

	c, err := publishservice.NewClient("Feed", client.WithHostPorts("127.0.0.1:6666"))
	if err != nil {
		panic(err)
	}
	publishClient = c
}


func PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp, err = publishClient.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil, errors.New("PublishAction Failed")
	}
	return resp, nil
}


func PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp, err = publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil, errors.New("PulishList Failed")
	}
	return resp, nil
}