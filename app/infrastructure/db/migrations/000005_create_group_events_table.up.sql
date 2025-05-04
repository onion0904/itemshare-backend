CREATE TABLE group_events (
    group_id VARCHAR(255) NOT NULL,
    event_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (group_id, event_id),
    CONSTRAINT fk_group_events_group FOREIGN KEY (group_id) REFERENCES `groups` (id) ON DELETE CASCADE,
    CONSTRAINT fk_group_events_event FOREIGN KEY (event_id) REFERENCES `events` (id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;