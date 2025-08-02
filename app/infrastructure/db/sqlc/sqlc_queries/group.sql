-- name: UpsertGroup :exec
INSERT INTO "groups" (
    id,
    name,
    created_at,
    updated_at
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(name),
    NOW(),
    NOW()
)
ON CONFLICT (id) DO UPDATE
SET
    name        = EXCLUDED.name,
    updated_at  = NOW();

-- name: DeleteGroup :exec
DELETE FROM "groups"
WHERE id = sqlc.arg(groupID);

-- name: FindGroup :one
SELECT
    id,
    name,
    created_at,
    updated_at
FROM "groups"
WHERE id = sqlc.arg(groupID);

-- name: AddUserToGroup :exec
INSERT INTO group_users (
    group_id,
    user_id
)
VALUES (
    sqlc.arg(groupID),
    sqlc.arg(userID)
)
ON CONFLICT (group_id, user_id) DO NOTHING;

-- name: AddEventToGroup :exec
INSERT INTO group_events (
    group_id,
    event_id
)
VALUES (
    sqlc.arg(groupID),
    sqlc.arg(eventID)
)
ON CONFLICT (group_id, event_id) DO NOTHING;

-- name: RemoveUserFromGroup :exec
DELETE FROM group_users
WHERE group_id = sqlc.arg(groupID)
    AND user_id  = sqlc.arg(userID);


-- name: GetUserIDsByGroupID :many
SELECT
    gu.user_id
FROM group_users AS gu
WHERE gu.group_id = sqlc.arg(groupID);

-- name: GetEventIDsByGroupID :many
SELECT
    ge.event_id
FROM group_events AS ge
WHERE ge.group_id = sqlc.arg(groupID);

