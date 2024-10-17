#!/bin/sh

# Install taskfile if not already installed
go install github.com/go-task/task/v3/cmd/task@latest

# Install golangci-lint if not already installed
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install goose if not already installed
go install github.com/pressly/goose/v3/cmd/goose@latest

# Install sqlc if not already installed
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

