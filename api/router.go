package api

import (
	"c-ademy/api/authn"
	"c-ademy/internal/config"
	"c-ademy/internal/db"
	"c-ademy/internal/db/sqlc"
	vmmanager "c-ademy/internal/vm_manager"
	"c-ademy/internal/vm_manager/language_mappings"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

func GetRouter(config *config.Environment) *echo.Echo {
	e := echo.New()

	// Configure logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	db, err := db.New(config.DbPath)
	if err != nil {
		e.Logger.Fatal(err)
	}

	authn := authn.New(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/authn/signup", func(c echo.Context) error {

		// Get the username from the request body
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err := c.Bind(&req)
		if err != nil {
			return err
		}

		// Check if the username is already taken

		exists, err := db.CheckUsernameExists(c.Request().Context(), req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		if exists == 1 {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "username already taken"})
			return nil
		}

		allowed, err := db.CheckUsernameAllowed(c.Request().Context(), req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		if allowed == 0 {
			// We don't want to give away whether the username is taken or not
			// so we return the same error message
			c.JSON(http.StatusBadRequest, map[string]string{"error": "username already taken"})
			return nil
		}

		// Hash the password
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		// Create a new user
		_, err = db.CreateUser(c.Request().Context(), sqlc.CreateUserParams{
			Username:     req.Username,
			PasswordHash: string(hash),
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		return nil
	})

	e.POST("/authn/get-token", func(c echo.Context) error {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		err := c.Bind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			return err
		}

		// Get the user from the database
		user, err := db.GetUserByUsername(c.Request().Context(), req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		// Check if the password is correct
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid username or password"})
			return nil
		}

		// Generate a token
		token, err := authn.GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		c.JSON(http.StatusOK, map[string]string{"token": token})
		return nil
	})

	e.POST("/exec", func(c echo.Context) error {
		var req struct {
			Language       string `json:"language"`
			FileName       string `json:"file_name"`
			Base64Code     string `json:"code"`
			Base64Input    string `json:"input"`
			ExpectedOutput string `json:"expected_output"`
			WebhookUrl     string `json:"webhook_url"`
		}

		err := c.Bind(&req)
		if err != nil {
			c.Logger().Errorf("Error binding request: %v", err)
			c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			return err
		}

		// Decode the base64 code
		decodedBytes, err := base64.StdEncoding.DecodeString(
			req.Base64Code,
		)
		if err != nil {
			c.Logger().Error("Error decoding string:", err)
			c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid base64 code"})
			return err
		}

		code := string(decodedBytes)

		// Decode the base64 input
		decodedBytes, err = base64.StdEncoding.DecodeString(
			req.Base64Input,
		)
		if err != nil {
			c.Logger().Error("Error decoding string:", err)
			c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid base64 input"})
			return err
		}

		input := string(decodedBytes)

		// Get the user from the database
		userID := c.Get("user_id").(int64)
		user, err := db.GetUserByID(c.Request().Context(), userID)
		if err != nil {
			c.Logger().Error("Error getting user:", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		// Validate the language
		language, err := language_mappings.ValidateLanguageKey(req.Language)
		if err != nil {
			c.Logger().Error("Error validating language:", err)
			c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid language"})
			return nil
		}

		// Create a new execution
		res, err := db.CreateExecution(c.Request().Context(), sqlc.CreateExecutionParams{
			UserID:         user.ID,
			Language:       req.Language,
			Code:           code,
			Input:          sql.NullString{String: input, Valid: input != ""},
			ExpectedOutput: sql.NullString{String: req.ExpectedOutput, Valid: req.ExpectedOutput != ""},
			WebhookUrl:     sql.NullString{String: req.WebhookUrl, Valid: req.WebhookUrl != ""},
		})

		if err != nil {
			c.Logger().Errorf("Error creating execution: %v", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		execId, err := res.LastInsertId()
		if err != nil {
			c.Logger().Errorf("Error getting last insert ID: %v", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		// Send the execution to the VM manager
		execRes, err := vmmanager.ExecuteProgram(c.Request().Context(), fmt.Sprintf("%d", execId), language, code, req.FileName)

		if err != nil {
			c.Logger().Errorf("Error executing program: %v", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		// Update the execution with the result
		err = db.UpdateExecution(c.Request().Context(), sqlc.UpdateExecutionParams{
			ID:     execId,
			Stdout: sql.NullString{String: execRes.Stdout, Valid: true},
			Stderr: sql.NullString{String: execRes.Stderr, Valid: true},
			ExitCode: sql.NullInt64{
				Int64: int64(execRes.ExitCode),
				Valid: true,
			},
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return err
		}

		return c.JSON(http.StatusOK, execRes)
	}, authn.AuthMiddleware())

	return e
}
