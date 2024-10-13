-- Allowlist table (unchanged)
CREATE TABLE allowlist (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    is_allowed BOOLEAN NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users table (unchanged)
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- API tokens table (unchanged)
CREATE TABLE api_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Executions table (updated)
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

-- Index for faster token lookups (unchanged)
CREATE INDEX idx_api_tokens_token ON api_tokens(token);

-- Trigger to update the updated_at timestamp for users (unchanged)
CREATE TRIGGER update_user_timestamp 
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- Trigger to ensure only allowlisted usernames can be used for registration (unchanged)
CREATE TRIGGER check_user_allowlist
BEFORE INSERT ON users
FOR EACH ROW
BEGIN
    SELECT CASE
        WHEN (SELECT COUNT(*) FROM allowlist WHERE username = NEW.username AND is_allowed = 1) = 0
        THEN RAISE(ABORT, 'Username not in allowlist or not allowed')
    END;
END;

-- Trigger to update the updated_at timestamp for executions
CREATE TRIGGER update_execution_timestamp 
AFTER UPDATE ON executions
FOR EACH ROW
BEGIN
    UPDATE executions SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;