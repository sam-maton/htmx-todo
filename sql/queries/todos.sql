-- name: GetTodosByUserId :many
SELECT id, created_at, updated_at, title, completed
FROM todos
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetTodoById :one
SELECT id, created_at, updated_at, title, completed, user_id
FROM todos
WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos(id, created_at, updated_at, user_id, title, completed)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
RETURNING id, created_at, updated_at, title, completed;

-- name: SetTodoCompleted :one
UPDATE todos SET completed = $2, updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, title, completed;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;