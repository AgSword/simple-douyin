package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse pack response
func SendResponse(c *app.RequestContext, response interface{}) {
	c.JSON(consts.StatusOK, response)
}

// 用户注册 handler 输入参数
type UserRegisterParam struct {
	UserName string `json:"username"` // 用户名
	PassWord string `json:"password"` // 用户密码
}

// 用户信息 输出参数
type UserParam struct {
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

type FavoriteActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	VideoId    int64  `json:"video_id,omitempty"`    // 视频id
	ActionType int32  `json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

//视频流输入参数
type Feedparam struct {
	Latest_time int64  `json:"latest_time,omitempty"` // Latest_time
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

//投稿接口输入参数
type PublishActionparam struct {
	Token string `json:"token"`
	Data  []byte `json:"data"`
	Title string `json:"title"`
}

type RelationActionParam struct {
	Token string `json:"token"`
	To_user_id  int `json:"to_user_id"`
	Action_type int `json:"action_type"`
}


type RelationFollowListParam struct{
	Token string `json:"token"`
	User_id int
}