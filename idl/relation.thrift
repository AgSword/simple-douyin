namespace go relation

struct User {
    1: i64 id
    2: string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: bool is_follow
}


struct RelationActionRequest {
    1: i64 user_id
    2: string token
    3: i64 to_user_id
    4: int32 action_type
}

struct RelationActionResponse {
    1: int32 status_code
    2: optional string status_msg
}

struct RelationFollowListRequest {
    1: i64 user_id
    2: string token
}

struct RelationFollowListResponse {
    1: int32 status_code
    2: optional status_msg
    3:  list<User> user_list
}

struct RelationFollowerListRequest {
    1: i64 user_id
    2: string token
}

struct RelationFollowerListResponse {
    1: int32 status_code
    2: string status_msg
    3: list<User> user_list
}

service RelationService {
    RelationActionResponse RelationAction (1: RelationActionRequest)
    RelationFollowListResponse RelationFollowList (1: RelationFollowerListRequest)
    RelationFollowerListResponse RelationFollowerList (1: RelationFollowerListRequest)
}
