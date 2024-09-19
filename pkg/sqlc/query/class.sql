-- name: GetClass :one
SELECT
  id,
  name,
  subject,
  grade
FROM classes
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;