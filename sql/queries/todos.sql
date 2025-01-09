-- name: GetTodosByUserId :many
SELECT id, created_at, updated_at, title, completed
FROM todos
WHERE user_id = $1;

-- name: CreateTodo :one
INSERT INTO todos(id, created_at, updated_at, user_id, title, completed)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
RETURNING id, created_at, updated_at, title, completed;