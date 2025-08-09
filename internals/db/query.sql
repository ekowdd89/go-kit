
-- name: FetchAuthors :many
SELECT id, name FROM author;

-- name: FetchAuthor :one
SELECT id, name FROM author WHERE id = $1;

-- name: CreateAuthor :one
INSERT INTO author (name) VALUES ($1) RETURNING *;

-- name: UpdateAuthor :one
UPDATE author SET name = $1 WHERE id = $2 RETURNING *;

-- name: DeleteAuthor :one
DELETE FROM author WHERE id = $1 RETURNING *;

-- name: FetchTodos :many
SELECT id, title, completed FROM todos;

-- name: FetchTodo :one
SELECT id, title, completed FROM todos WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING *;

-- name: UpdateTodo :one
UPDATE todos SET title = $1, completed = $2 WHERE id = $3 RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos WHERE id = $1 RETURNING *;