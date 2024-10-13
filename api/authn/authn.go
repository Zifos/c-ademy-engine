package authn

import (
	"c-ademy/internal/db"
	"c-ademy/internal/db/sqlc"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"database/sql"

	"github.com/labstack/echo/v4"
)

type Authn struct {
	db *db.Queries
}

func New(db *db.Queries) *Authn {
	return &Authn{
		db: db,
	}
}

func (authn *Authn) GenerateToken(userId int64) (string, error) {
	// Generate a random token
	tokenBytes := make([]byte, 32) // 256 bits
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	// Set expiration time (e.g., 24 hours from now)
	expiresAt := sql.NullTime{
		Time:  time.Now().Add(24 * time.Hour),
		Valid: true,
	}

	// Create token in the database
	_, err = authn.db.CreateToken(context.Background(), sqlc.CreateTokenParams{
		UserID:    userId,
		Token:     token,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (authn *Authn) ExtractBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid Authorization header format")
	}

	token := parts[1]
	if token == "" {
		return "", errors.New("token is missing")
	}

	return token, nil
}

func (authn *Authn) AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := authn.ExtractBearerToken(c.Request())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			}

			// Validate the token
			valid, err := authn.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error validating token"})
			}

			if !valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			// Optionally, you can add user information to the context
			userID, err := authn.GetUserIDFromToken(token)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting user information"})
			}
			c.Set("user_id", userID)

			// Token is valid, call the next handler
			return next(c)
		}
	}
}

func (authn *Authn) GetUserIDFromToken(token string) (int64, error) {
	// Fetch the token from the database
	dbToken, err := authn.db.GetToken(context.Background(), token)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("token not found")
		}
		return 0, err
	}

	return dbToken.UserID, nil
}

func (authn *Authn) ValidateToken(token string) (bool, error) {
	// Fetch the token from the database
	dbToken, err := authn.db.GetToken(context.Background(), token)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // Token not found
		}
		return false, err // Database error
	}

	// Check if the token has expired
	if dbToken.ExpiresAt.Valid && dbToken.ExpiresAt.Time.Before(time.Now()) {
		return false, nil // Token has expired
	}

	return true, nil // Token is valid
}
