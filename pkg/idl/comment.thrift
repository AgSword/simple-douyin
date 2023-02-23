namespace go comment

struct User {
    1: i64 id
    2: string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: bool is_follow
}

struct Comment {
    1: i64 id
    2: User 
}

struct CommentActionRequest {
    1: i64 user_id
    2: string token
    3: i64 video_id
    4: int32 action_type
    5: optional string comment_text
    6: optional i64 comment_id
}

struct CommentActionResponse {
    1: int32 status_code
    2: optional string status_msg
    3: optional Comment comment
}

struct CommentListRequest {
    1: string token
    2: i64 video_id
}

struct CommentListResponse {
    1: int32 status_code
    2: optional string status_msg
    3: list<Comment> comment_list
}

service CommentService {
    CommentActionResponse CommentAction (1: CommentActionRequest)
    CommentListResponse CommentList (1: CommentListRequest)
}

