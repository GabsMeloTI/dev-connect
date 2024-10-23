CREATE TABLE "Post" (
                        id SERIAL PRIMARY KEY,
                        user_id INT REFERENCES "User"(id) ON DELETE CASCADE,
                        content VARCHAR NOT NULL,
                        image_url VARCHAR,
                        likes INT DEFAULT 0,
                        shares INT DEFAULT 0,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);