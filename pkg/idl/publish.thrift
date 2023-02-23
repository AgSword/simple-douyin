namespace go publish

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

struct PublishActionRequest {
    1: string token
    2: bytes data
    3: string title
}

struct PublishActionResponse {
    1: int32 status_code
    2: optional string status_msg
}

struct PublishListRequest {
    1: i64 user_id
    2: string token
}

struct PublishListResponse {
    1: int32 status_code
    2: optional string status_msg
    3: list<Video> video_list
}

service PublishService {
    PublishActionResponse PublishAction (1: PublishActionRequest)
    PublishListResponse PublishList (1: PublishListRequest)
}


