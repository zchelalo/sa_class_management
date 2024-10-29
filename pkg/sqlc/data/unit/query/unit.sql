-- name: GetUnit :one
SELECT
  id,
  name
FROM units
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;