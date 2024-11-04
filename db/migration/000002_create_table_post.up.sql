CREATE TABLE "Post" (
                        id BIGSERIAL PRIMARY KEY,
                        user_id BIGINT REFERENCES "User"(id) ON DELETE CASCADE NOT NULL,
                        content VARCHAR NOT NULL,
                        image_url VARCHAR,
                        likes INT DEFAULT 0,
                        shares INT DEFAULT 0,
                        archive BOOL NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);