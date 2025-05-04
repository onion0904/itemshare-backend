DROP TABLE IF EXISTS user_events (
    user_id VARCHAR(255) NOT NULL,
    event_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_id, event_id),
    CONSTRAINT fk_user_events_user FOREIGN KEY (user_id) REFERENCES `users` (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_events_event FOREIGN KEY (event_id) REFERENCES `events` (id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;