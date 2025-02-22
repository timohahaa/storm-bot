-- name: GetUser :one
SELECT * FROM core.user
WHERE telegram_id = $1
    AND deleted_at IS NULL;

-- name: CreateUser :one
INSERT INTO core.user (
    telegram_id
    , is_admin
) VALUES (
    $1, $2
)
RETURNING *
;

-- name: CreateLink :one
INSERT INTO core.link (
    chat_id 
    , link 
) VALUES (
    $1, $2
)
RETURNING *
;

-- name: MonthLinkStats :many
SELECT * FROM core.link
WHERE EXTRACT(MONTH FROM created_at) = $1
    AND deleted_at IS NULL
;
