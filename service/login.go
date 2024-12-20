package service

import (
	"os"
	"time"

	"github.com/ahay12/api-test/database"
	"github.com/ahay12/api-test/helper"
	"github.com/ahay12/api-test/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = os.Getenv("API_KEY")

func Login(ctx *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := ctx.BodyParser(&input); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil, err.Error())
		return err
	}

	var user model.Users

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			helper.RespondJSON(ctx, fiber.StatusNotFound, "User not found", nil, err.Error())
		} else {
			// Some other error occurred
			helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Something went wrong", nil, err.Error())
		}
		return err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Unauthorized", nil, err.Error())
		return err
	}

	token, err := generateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to generate JWT", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "token", token, nil)
	return nil

}

func generateJWT(id uint, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID": id,
		"email":  email,
		"role":   role,
		"exp":    time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
