

CREATE TABLE chat_members (
    chat_id   UUID NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    user_id   UUID NOT NULL,
    joined_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (chat_id, user_id)
);

CREATE INDEX idx_chat_members_user_id ON chat_members(user_id);
