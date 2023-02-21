package rpc

import(
	"github.com/AgSword/simpleDouyin/kitex_gen/feed/feedservice"
	"github.com/cloudwego/kitex/client"
	"github.com/AgSword/simpleDouyin/kitex_gen/feed"
	"context"
	"errors"
)

var feedClient feedservice.Client

// feed RPC 客户端初始化
func init() {

	c, err := feedservice.NewClient("Feed", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
	feedClient = c
}


func GetUserFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	resp, err = feedClient.GetUserFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil, errors.New("Action Failed")
	}
	return resp, nil
}
