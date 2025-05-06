CREATE TABLE group_users (
    group_id VARCHAR(255) NOT NULL REFERENCES "groups"(id) ON DELETE CASCADE,
    user_id VARCHAR(255) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (group_id, user_id)
);