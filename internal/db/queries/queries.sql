-- User-related queries
-- name: GetUserByID :one
SELECT id, username, password_hash, created_at, updated_at
FROM users
WHERE id = ?;

-- name: GetUserByUsername :one
SELECT id, username, password_hash, created_at, updated_at
FROM users
WHERE username = ?;

-- name: CreateUser :execresult
INSERT INTO users (username, password_hash)
VALUES (?, ?);

-- name: UpdateUsername :exec
UPDATE users
SET username = ?
WHERE id = ?;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = ?
WHERE id = ?;

-- name: CheckUsernameExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE username = ?) AS username_exists;

-- Token-related queries
-- name: GetToken :one
SELECT id, user_id, token, expires_at, created_at
FROM api_tokens
WHERE token = ?;

-- name: CreateToken :execresult
INSERT INTO api_tokens (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: DeleteToken :exec
DELETE FROM api_tokens
WHERE token = ?;

-- Allowlist-related queries
-- name: CheckUsernameAllowed :one
SELECT EXISTS(SELECT 1 FROM allowlist WHERE username = ? AND is_allowed = 1) AS username_allowed;

-- name: AddToAllowlist :exec
INSERT INTO allowlist (username, is_allowed)
VALUES (?, ?)
ON CONFLICT(username) DO UPDATE SET is_allowed = excluded.is_allowed;

-- name: RemoveFromAllowlist :exec
DELETE FROM allowlist
WHERE username = ?;

-- name: UpdateAllowlistStatus :exec
UPDATE allowlist
SET is_allowed = ?
WHERE username = ?;

-- Execution-related queries
-- name: CreateExecution :execresult
INSERT INTO executions (
    user_id, language, code, input, expected_output, webhook_url
) VALUES (?, ?, ?, ?, ?, ?);

-- name: GetExecutionByID :one
SELECT id, user_id, language, code, input, expected_output, stdout, stderr, webhook_url, created_at, updated_at
FROM executions
WHERE id = ?;

-- name: UpdateExecutionOutput :exec
UPDATE executions
SET stdout = ?
WHERE id = ?;

-- name: ListUserExecutions :many
SELECT id, language, code, input, expected_output, stdout, stderr, webhook_url, created_at, updated_at
FROM executions
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ?;

-- name: UpdateExecution :exec
UPDATE executions
SET 
    stdout = ?,
    stderr = ?,
    exit_code = ?
WHERE id = ?;