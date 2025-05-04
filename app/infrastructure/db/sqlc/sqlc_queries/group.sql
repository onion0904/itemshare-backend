-- name: UpsertGroup :exec
INSERT INTO
    `groups` (
        id,
        name,
        icon,
        created_at,
        updated_at
    )
VALUES
    (
        sqlc.arg(id),
        sqlc.arg(name),
        sqlc.arg(icon),
        NOW(),
        NOW()
    ) ON DUPLICATE KEY
UPDATE
    name = sqlc.arg(name),
    icon = sqlc.arg(icon),
    updated_at = NOW();

-- name: DeleteGroup :exec
DELETE FROM
    `groups`
WHERE
    id = sqlc.arg(groupID);

-- name: FindGroup :one
SELECT
    id,
    name,
    icon,
    created_at,
    updated_at
FROM
    `groups`
WHERE
    id = sqlc.arg(groupID);

-- name: AddUserToGroup :exec
INSERT INTO
    group_users (group_id, user_id)
VALUES
    (
        sqlc.arg(groupID),
        sqlc.arg(userID)
    )
ON DUPLICATE KEY UPDATE
    group_id = sqlc.arg(groupID), 
    user_id = sqlc.arg(userID);

-- name: AddEventToGroup :exec
INSERT INTO
    group_events (group_id, event_id)
VALUES
    (
        sqlc.arg(groupID),
        sqlc.arg(eventID)
    )
ON DUPLICATE KEY UPDATE
    group_id = sqlc.arg(groupID), 
    event_id = sqlc.arg(eventID);

-- name: RemoveUserFromGroup :exec
DELETE FROM 
    group_users
WHERE 
    group_id = ? AND user_id = ?;


-- name: GetUserIDsByGroupID :many
SELECT
    gu.user_id
FROM
    group_users gu
WHERE
    gu.group_id = sqlc.arg(groupID);

-- name: GetEventIDsByGroupID :many
SELECT
    ge.event_id
FROM
    group_events ge
WHERE
    ge.group_id = sqlc.arg(groupID);

