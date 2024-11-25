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

-- name: GetClassByCode :one
SELECT
  id,
  name,
  subject,
  grade
FROM classes
WHERE code = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: ListClasses :many
SELECT
  classes.id,
  classes.name,
  classes.subject,
  classes.grade
FROM classes
INNER JOIN members ON classes.id = members.class_id
WHERE members.user_id = $1
AND classes.deleted_at IS NULL
AND members.deleted_at IS NULL
ORDER BY classes.created_at DESC
OFFSET $2
LIMIT $3;

-- name: CountClasses :one
SELECT COUNT(*)
FROM classes
INNER JOIN members ON classes.id = members.class_id
WHERE members.user_id = $1
AND classes.deleted_at IS NULL
AND members.deleted_at IS NULL;

-- name: CreateClass :one
INSERT INTO classes (
  id,
  name,
  subject,
  grade,
  code
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING *;