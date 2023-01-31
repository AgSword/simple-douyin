namespace go user

struct User {
    1: i64 id
    2: string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: bool is_follow
}

struct UserRegisterRequest {
    1: string username ( vt.min_size = "1", vt.max_size = "32")
    2: string password ( vt.min_size = "1", vt.max_size = "32")
}

struct UserRegisterResponse {
    1: i32 status_code; 
    2: optional string status_msg; 
    3: i64 user_id; 
    4: string token; 
}

struct UserLoginRequest {
    1: string username (vt.min_size = "1", vt.max_size = "32")
    2: string password (vt.min_size = "1", vt.max_size = "32")
}

struct UserLoginResponse {
    1: int32 status_code; 
    2: optional string status_msg; 
    3: int64 user_id; 
    4: string token; 
}

struct UserRequest{
    1: int64 user_id; 
    2: string token; 
}

struct UserResponse{
    1: int32 status_code;
    2: optional string status_msg; 
    3: User user;
}

service UserService {
    UserRegisterResponse Register (1: UserRegisterRequest req)
    UserLoginResponse Login (1: UserLoginRequest req)
    UserResponse GetUserById (1: UserRequest req)
}