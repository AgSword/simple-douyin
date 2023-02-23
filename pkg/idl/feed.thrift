namespace go feed

struct User {
    1: i64 id
    2: string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: bool is_follow
}

struct Video {
    1: i64 id
    2: User author
    3: string play_url
    4: string cover_url
    5: i64 favorite_count
    6: i64 comment_count
    7: bool is_favorite
    8: string title
}

struct FeedRequest {
    1: optional i64 latest_time
    2: optional string token
}

struct FeedResponse {
    1: int32 status_code
    2: optional string status_msg
    3: list<Video> video_list
    4: optional i64 next_time
}

struct VideoIdRequest {
    1: i64 video_id
    2: i64 search_id
}

service FeedService {
    FeedResponse GetUserFeed (1: FeedRequest)
    Video GetVideoById (1: VideoIdRequest)
}
