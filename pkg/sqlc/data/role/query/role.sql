-- name: GetRole :one
SELECT
  id,
  key
FROM roles
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;