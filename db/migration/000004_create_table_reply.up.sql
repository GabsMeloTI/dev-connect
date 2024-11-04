CREATE TABLE "Reply" (
                         id BIGSERIAL PRIMARY KEY,
                         user_id BIGINT REFERENCES "User"(id) ON DELETE CASCADE,
                         comment_id BIGINT REFERENCES "Comment"(id) ON DELETE CASCADE,
                         content VARCHAR NOT NULL,
                         likes INT DEFAULT 0,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);