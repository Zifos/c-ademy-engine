-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- Allowlist table
CREATE TABLE allowlist (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    is_allowed BOOLEAN NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- API tokens table
CREATE TABLE api_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Executions table
CREATE TABLE executions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    code TEXT NOT NULL,
    input TEXT,
    expected_output TEXT,
    stdout TEXT,
    stderr TEXT,
    exit_code INTEGER,
    webhook_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Index for faster token lookups
CREATE INDEX idx_api_tokens_token ON api_tokens(token);

-- Trigger to update the updated_at timestamp for users
-- +goose StatementBegin
CREATE TRIGGER update_user_timestamp 
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- Trigger to ensure only allowlisted usernames can be used for registration
-- +goose StatementBegin
CREATE TRIGGER check_user_allowlist
BEFORE INSERT ON users
FOR EACH ROW
BEGIN
    SELECT CASE
        WHEN (SELECT COUNT(*) FROM allowlist WHERE username = NEW.username AND is_allowed = 1) = 0
        THEN RAISE(ABORT, 'Username not in allowlist or not allowed')
    END;
END;
-- +goose StatementEnd

-- Trigger to update the updated_at timestamp for executions
-- +goose StatementBegin
CREATE TRIGGER update_execution_timestamp 
AFTER UPDATE ON executions
FOR EACH ROW
BEGIN
    UPDATE executions SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TRIGGER IF EXISTS update_execution_timestamp;
DROP TRIGGER IF EXISTS check_user_allowlist;
DROP TRIGGER IF EXISTS update_user_timestamp;
DROP INDEX IF EXISTS idx_api_tokens_token;
DROP TABLE IF EXISTS executions;
DROP TABLE IF EXISTS api_tokens;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS allowlist;