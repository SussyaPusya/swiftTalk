
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE chats (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type       VARCHAR(10) NOT NULL CHECK (type IN ('private', 'group')),
    name       VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
