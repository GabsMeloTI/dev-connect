CREATE TABLE "Follower" (
                            id SERIAL PRIMARY KEY,
                            follower_id INT REFERENCES "User"(id) ON DELETE CASCADE,
                            followed_id INT REFERENCES "User"(id) ON DELETE CASCADE,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);