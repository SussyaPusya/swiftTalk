package models

type CreateChatDTO struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AddMemberDTO struct {
	UserID string `json:"user_id"`
}
type GetChatMembersDTO struct {
	Members []Member `json:"members"`
}
type SendMessageDTO struct {
	SenderID string `json:"sender_id"`
	Content  string `json:"content"`
}

type GetChatsDTO struct {
	Chats []Chat `json:"chats"`
}

type GetMessagesDTO struct {
	Messages []Message `json:"messages"`
}

type DeleteMessageDTO struct {
	MessageID string `json:"message_id"`
}

type DeleteMemberFromChatDTO struct {
	UserID string `json:"user_id"`
}
