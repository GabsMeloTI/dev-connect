CREATE TABLE "Comment" (
                           id SERIAL PRIMARY KEY,
                           user_id INT REFERENCES "User"(id) ON DELETE CASCADE,
                           post_id INT REFERENCES "Post"(id) ON DELETE CASCADE,
                           content VARCHAR NOT NULL,
                           likes INT DEFAULT 0,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
