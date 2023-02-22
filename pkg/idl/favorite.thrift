namespace go favorite

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

struct FavoriteActionRequest {
    1: i64 user_id
    2: string token
    3: i64 video_id
    4: action_type
}

struct FavoriteActionResponse {
    1: int32 status_code
    2: optional string status_msg
}

struct FavoriteListRequest {
    1: i64 user_id
    2: string token
}

struct FavoriteListResponse {
    1: int32 status_code
    2: optional string status_msg
    3: list<Video> video_list
}

service FavoriteService {
    FavoriteActionResponse FavoriteAction (1: FavoriteListRequest)
    FavoriteListResponse FavoriteList (1: FavoriteListRequest)
}

