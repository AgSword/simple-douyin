namespace go favorite

struct douyin_favorite_action_request {
    //只有1和文档中的不一致，让其他服务鉴权，这里只收user_id
    //1: string  token // 用户鉴权
    1: i64  user_id // 用户id
    2: i64  video_id // 视频id
    3: i32 action_type // 1-点赞，2-取消点赞

}

struct douyin_favorite_action_response  {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

struct douyin_favorite_list_request {
    1: i64 user_id // 用户id
    // 让其他服务鉴权，这里只收user_id
    // 2: string token  // 用户鉴权token
}

struct douyin_favorite_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: list<Video> video_list // 用户点赞视频列表
}

struct Video {
    1: i64 id  // 视频唯一标识
    2: User author // 视频作者信息
    3: string play_url  // 视频播放地址
    4: string cover_url // 视频封面地址
    5: i64 favorite_count  // 视频的点赞总数
    6: i64 comment_count  // 视频的评论总数
    7: bool is_favorite  // true-已点赞，false-未点赞
    8: string title  // 视频标题
}

struct User {
    1: i64 id  // 用户id
    2: string name  // 用户名称
    3: optional i64 follow_count  // 关注总数
    4: optional i64 follower_count  // 粉丝总数
    5: bool is_follow  // true-已关注，false-未关注
}

service Favorite {
    douyin_favorite_action_response Action(1: douyin_favorite_action_request req)
    douyin_favorite_list_response List(2: douyin_favorite_list_request req)
}