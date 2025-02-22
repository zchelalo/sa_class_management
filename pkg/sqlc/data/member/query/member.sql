-- name: GetMember :one
SELECT
  members.id,
  members.user_id,
  members.class_id,
  roles.id as role_id,
  roles.key as role_key
FROM members
INNER JOIN roles ON members.role_id = roles.id
WHERE members.user_id = $1
AND members.class_id = $2
AND members.deleted_at IS NULL
AND roles.deleted_at IS NULL
LIMIT 1;

-- name: ListMembers :many
SELECT
  members.id,
  members.user_id,
  members.class_id,
  roles.id as role_id,
  roles.key as role_key
FROM members
INNER JOIN roles ON members.role_id = roles.id
WHERE members.class_id = $1
AND members.deleted_at IS NULL
AND roles.deleted_at IS NULL
ORDER BY members.created_at DESC
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