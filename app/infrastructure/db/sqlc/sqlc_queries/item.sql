-- itemの追加
-- name: InsertItem :exec
INSERT INTO items (
    id,
    group_id,
    name,
    created_at,
    updated_at
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(group_id),
    sqlc.arg(name),
    NOW(),
    NOW()
);

-- itemの削除
-- name: DeleteItem :exec
DELETE FROM items
WHERE id = sqlc.arg(itemID);

-- itemIDからitemを取得するクエリ
-- name: GetItemByID :one
SELECT *
FROM items
WHERE id = sqlc.arg(item_id);

-- groupID から item をすべて取得するクエリ
-- name: GetItemsByGroupID :many
SELECT *
FROM items
WHERE group_id = sqlc.arg(group_id)
ORDER BY created_at ASC;
