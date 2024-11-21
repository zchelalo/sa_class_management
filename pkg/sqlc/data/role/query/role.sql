-- name: GetRoleByID :one
SELECT
  id,
  key
FROM roles
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetRoleByKey :one
SELECT
  id,
  key
FROM roles
WHERE key = $1
AND deleted_at IS NULL
LIMIT 1;