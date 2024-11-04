CREATE TABLE "Follower" (
                            id BIGSERIAL PRIMARY KEY,
                            follower_id BIGINT REFERENCES "User"(id) ON DELETE CASCADE,
                            followed_id BIGINT REFERENCES "User"(id) ON DELETE CASCADE,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);