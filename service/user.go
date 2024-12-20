package service

import (
	"errors"

	"github.com/ahay12/api-test/database"
	"github.com/ahay12/api-test/helper"
	"github.com/ahay12/api-test/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(ctx *fiber.Ctx) error {
	var users []model.Users
	DB := database.DB
	if err := DB.Find(&users).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to fetch users", nil, err.Error())
		return err
	}

	helper.RespondJSON(ctx, fiber.StatusOK, "User fetched successfully", users, nil)
	return nil
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user model.Users
	if err := database.DB.First(&user, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "user not found", nil, err.Error())
		return err
	}

	helper.RespondJSON(ctx, fiber.StatusOK, "User created success", user, nil)
	return nil
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(model.Users)
	if err := ctx.BodyParser(user); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "cannot pasrse JSON", nil, err.Error())
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "failed to hash password", nil, err.Error())
		return err
	}

	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to create user", nil, []helper.ErrorField{
			{
				ID:      "database",
				Value:   "creation",
				Caused:  "database error",
				Message: err.Error(),
			},
		})
		return err
	}

	helper.RespondJSON(ctx, fiber.StatusAccepted, "User created success", fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}, nil)
	return nil
}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user model.Users
	if err := database.DB.First(&user, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "User not found", nil, err.Error())
		return err
	}
	if err := ctx.BodyParser(&user); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "cannot parse JSON", nil, err.Error())
		return err
	}
	if err := database.DB.Save(&user).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to update user", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "Success update user", user, nil)
	return nil
}

func DeleteUser(ctx *fiber.Ctx) error {
	if database.DB == nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Database connection not established", nil, nil)
		return errors.New("database connection not established")
	}

	id := ctx.Params("id")
	var user model.Users

	if err := database.DB.First(&user, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "User not found", nil, err.Error())
		return err
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to delete user", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "User deleted", nil, nil)
	return nil
}
