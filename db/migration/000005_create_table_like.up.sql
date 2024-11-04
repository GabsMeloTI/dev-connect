CREATE TABLE "Like" (
                        id BIGSERIAL PRIMARY KEY,
                        user_id BIGINT REFERENCES "User"(id) ON DELETE CASCADE,
                        post_id BIGINT REFERENCES "Post"(id) ON DELETE CASCADE,
                        comment_id BIGINT REFERENCES "Comment"(id) ON DELETE CASCADE,
                        reply_id BIGINT REFERENCES "Reply"(id) ON DELETE CASCADE,
                        type VARCHAR(20) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
