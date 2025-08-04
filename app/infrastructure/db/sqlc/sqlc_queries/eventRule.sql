-- eventRuleの追加と更新
-- name: UpsertEventRule :exec
INSERT INTO event_rules (
    user_id,
    item_id,
    normal_limit,
    important_limit
)
VALUES (
    sqlc.arg(user_id),
    sqlc.arg(item_id),
    sqlc.arg(normal_limit),
    sqlc.arg(important_limit)
)
ON CONFLICT (user_id, item_id) DO UPDATE
SET
    normal_limit = EXCLUDED.normal_limit,
    important_limit = EXCLUDED.important_limit;

-- userIDとitemIDからeventRuleの取得
-- name: GetEventRuleByUserAndItem :one
SELECT
    normal_limit,
    important_limit
FROM event_rules
WHERE user_id = sqlc.arg(user_id)
  AND item_id = sqlc.arg(item_id);

-- name: GetEventRulesByItemID :many
SELECT *
FROM event_rules
WHERE item_id = sqlc.arg(item_id)
ORDER BY user_id;