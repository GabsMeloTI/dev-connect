CREATE TABLE "Reply" (
                         id SERIAL PRIMARY KEY,
                         user_id INT REFERENCES "User"(id) ON DELETE CASCADE,
                         comment_id INT REFERENCES "Comment"(id) ON DELETE CASCADE,
                         content VARCHAR NOT NULL,
                         likes INT DEFAULT 0,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);