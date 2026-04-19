package models

type Chat struct {
	ChatID    string `json:"chat_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}

type Message struct {
	MessageID string `json:"message_id"`
	ChatID    string `json:"chat_id"`
	SenderID  string `json:"sender_id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type Member struct {
	UserID   string `json:"user_id"`
	JoinedAt string `json:"joined_at"`
}
