namespace go message

struct Message {
    1: i64 id
    2: string content
    3: string create_time
}

struct MesssageChatRequest {
    1: string token
    2: i64 to_user_id
}

struct MessageChatResponse {
    1: int32 status_code
    2: string status_msg
    3: list<Message> message_list
}

struct RelationActionRequest {
    1: string token
    2: i64 to_user_id
    3: int32 action_type
    4: string content
}

struct RelationActionResponse {
    1: int32 status_code
    2: string status_msg
}

service MessageService {
    MessageChatResponse GetChatMessage (1: MesssageChatRequest)
    RelationActionResponse PostMessage (1: RelationActionRequest)
}