package middleware

import (
	"os"
	"strings"

	"github.com/ahay12/api-test/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("API_KEY")

// AdminMiddleware checks if the user has an admin role
func AdminMiddleware(ctx *fiber.Ctx) error {
	// Get the JWT from the Authorization header
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Missing or invalid token", nil, nil)
		return nil
	}

	// Remove "Bearer " prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the JWT
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Unauthorized", nil, err.Error())
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Invalid token", nil, nil)
		return nil
	}

	// Check if the role is admin
	if role, roleExists := claims["role"].(string); !roleExists || role != "admin" {
		helper.RespondJSON(ctx, fiber.StatusForbidden, "Forbidden, only admins can access", nil, nil)
		return nil
	}

	// Extract userID from claims
	userID, userIDExists := claims["userID"].(float64) // JWT claims store numbers as float64
	if !userIDExists {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "User ID is required Middleware", nil, nil)
		// Stop execution and send unauthorized response
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// Store userID in the context for use in later handlers
	ctx.Locals("userID", uint(userID)) // Convert float64 to uint if needed

	return ctx.Next()
}

func UserMiddleware(ctx *fiber.Ctx) error {
	// Get the JWT from the Authorization header
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Missing or invalid token", nil, nil)
		return nil
	}

	// Remove "Bearer " prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the JWT
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Unauthorized", nil, err.Error())
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Invalid token", nil, nil)
		return nil
	}

	// Check if the role is admin
	// if role, roleExists := claims["role"].(string); !roleExists || role != "admin" {
	// 	helper.RespondJSON(ctx, fiber.StatusForbidden, "Forbidden, only admins can access", nil, nil)
	// 	return nil
	// }

	// Extract userID from claims
	userID, userIDExists := claims["userID"].(float64) // JWT claims store numbers as float64
	if !userIDExists {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "User ID is required Middleware", nil, nil)
		// Stop execution and send unauthorized response
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// Store userID in the context for use in later handlers
	ctx.Locals("userID", uint(userID)) // Convert float64 to uint if needed

	return ctx.Next()
}
