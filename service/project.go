package service

import (
	"errors"
	"github.com/ahay12/api-test/database"
	"github.com/ahay12/api-test/helper"
	"github.com/ahay12/api-test/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetProjects(ctx *fiber.Ctx) error {
	var articles []model.Project
	DB := database.DB
	if err := DB.Find(&articles).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to fetch articles", nil, err.Error())
		return err
	}

	helper.RespondJSON(ctx, fiber.StatusOK, "Articles fetched successfully", articles, nil)
	return nil
}

func GetProject(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var project model.Project

	// Start a database transaction to ensure atomicity
	tx := database.DB.Begin()

	// Fetch the article including its associated likes
	if err := database.DB.Preload("Likes").First(&project, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "Article not found", nil, err.Error())
		return err
	}

	// Increment the view count
	if err := tx.Model(&project).Update("view", project.View+1).Error; err != nil {
		tx.Rollback()
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to update view count", nil, err.Error())
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to commit transaction", nil, err.Error())
		return err
	}

	articleData := fiber.Map{
		"ID":        project.ID,
		"CreatedAt": project.CreatedAt,
		"UpdatedAt": project.UpdatedAt,
		"DeletedAt": project.DeletedAt,
		"image":     project.Image,
		"title":     project.Title,
		"Goals":     project.Goals,
		"Fund":      project.Fund,
		"tag":       project.Tag,
		"view":      project.View,
	}

	// Respond with the article data and like count
	helper.RespondJSON(ctx, fiber.StatusOK, "Success get article", fiber.Map{
		"article": articleData,
	}, nil)

	return nil
}

func CreateProject(ctx *fiber.Ctx) error {
	article := new(model.Project)
	if err := ctx.BodyParser(article); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil, err.Error())
		return err
	}

	if err := database.DB.Create(&article).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to create article", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "Successfully create article", article, nil)
	return nil
}

func UpdateProject(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var article model.Project
	if err := database.DB.First(&article, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "Article not found", nil, err.Error())
		return err
	}

	if err := ctx.BodyParser(&article); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil, err.Error())
		return err
	}
	if err := database.DB.Save(&article).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to update article", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "Successfully update article", article, nil)
	return nil
}

func DeleteProject(ctx *fiber.Ctx) error {
	if database.DB == nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Database connection not established", nil, nil)
		return errors.New("database connection not established")
	}
	log.Println("DeleteArticle function called")
	id := ctx.Params("id")
	log.Println("Delete request received for article ID:", id) // Tambahkan logging

	var article model.Project
	if err := database.DB.First(&article, id).Error; err != nil {
		log.Println("Article not found:", err)
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Article not found", nil, err.Error())
		return err
	}

	if err := database.DB.Delete(&article).Error; err != nil {
		log.Println("Failed to delete article:", err)
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to delete article", nil, err.Error())
		return err
	}
	log.Println("Article deleted successfully")
	helper.RespondJSON(ctx, fiber.StatusOK, "Article deleted successfully", nil, nil)
	return nil
}
