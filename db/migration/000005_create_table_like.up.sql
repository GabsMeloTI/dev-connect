CREATE TABLE "Like" (
                        id SERIAL PRIMARY KEY,
                        user_id INT REFERENCES "User"(id) ON DELETE CASCADE,
                        post_id INT REFERENCES "Post"(id) ON DELETE CASCADE,
                        comment_id INT REFERENCES "Comment"(id) ON DELETE CASCADE,
                        reply_id INT REFERENCES "Reply"(id) ON DELETE CASCADE,
                        type VARCHAR(20) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
