CREATE TABLE group_events (
    group_id VARCHAR(255) NOT NULL REFERENCES "groups"(id) ON DELETE CASCADE,
    event_id VARCHAR(255) NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    PRIMARY KEY (group_id, event_id)
);