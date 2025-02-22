-- name: GetMember :one
SELECT
  id,
  role_id,
  user_id,
  class_id
FROM members
WHERE user_id = $1
AND class_id = $2
AND deleted_at IS NULL
LIMIT 1;

-- name: ListMembers :many
SELECT
  id,
  role_id,
  user_id,
  class_id
FROM members
WHERE class_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: CountMembers :one
SELECT count(*)
FROM members
WHERE class_id = $1
AND deleted_at IS NULL;

-- name: CreateMember :one
INSERT INTO members (
  id,
  role_id,
  user_id,
  class_id
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING *;

-- name: DeleteMember :exec
UPDATE members
SET deleted_at = now()
WHERE user_id = $1
AND class_id = $2
AND deleted_at IS NULL;