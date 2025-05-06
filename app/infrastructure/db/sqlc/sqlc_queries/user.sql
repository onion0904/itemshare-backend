-- name: UpsertUser :exec
INSERT INTO users (
    id,
    last_name,
    first_name,
    email,
    password,
    icon,
    created_at,
    updated_at
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(last_name),
    sqlc.arg(first_name),
    sqlc.arg(email),
    sqlc.arg(password),
    sqlc.arg(icon),
    NOW(),
    NOW()
)
ON CONFLICT (id) DO UPDATE
SET
    last_name   = EXCLUDED.last_name,
    first_name  = EXCLUDED.first_name,
    email       = EXCLUDED.email,
    password    = EXCLUDED.password,
    icon        = EXCLUDED.icon,
    updated_at  = NOW();  


-- name: FindUserByID :one
SELECT
    id,
    last_name,
    first_name,
    email,
    password,
    icon,
    created_at,
    updated_at
FROM users
WHERE id = sqlc.arg(id);


-- name: FindUserByEmailPassword :one
SELECT
    id,
    last_name,
    first_name,
    email,
    password,
    icon,
    created_at,
    updated_at
FROM users
WHERE email = sqlc.arg(email)
    AND password = sqlc.arg(password);


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = sqlc.arg(id);


-- name: ExistUser :one
SELECT EXISTS(
    SELECT 1
    FROM users
    WHERE email = sqlc.arg(email)
        AND password = sqlc.arg(password)
) AS exists_user;


-- name: GetGroupIDsByUserID :many
SELECT gu.group_id
FROM group_users AS gu
WHERE gu.user_id = sqlc.arg(userID);


-- name: GetEventIDsByUserID :many
SELECT ue.event_id
FROM user_events AS ue
WHERE ue.user_id = sqlc.arg(userID);
