CREATE TABLE event_rules (
    user_id VARCHAR(255) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    item_id VARCHAR(255) NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    normal_limit INT NOT NULL,
    important_limit INT NOT NULL,
    PRIMARY KEY (user_id, item_id)
);