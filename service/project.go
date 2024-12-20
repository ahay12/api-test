package service

import (
	"errors"
	"fmt"
	"github.com/ahay12/api-test/database"
	"github.com/ahay12/api-test/helper"
	"github.com/ahay12/api-test/model"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
)

func GetProjects(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	sort := ctx.Query("sort", "id desc")
	order := ctx.Query("order", "asc")
	filterName := ctx.Query("title", "")

	offset := (page - 1) * limit

	var project []model.Project
	DB := database.DB
	if err := DB.Find(&project).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to fetch Project", nil, err.Error())
		return err
	}

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	if filterName != "" {
		DB = DB.Where("title LIKE ?", filterName)
	}

	DB = DB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := DB.Limit(limit).Offset(offset).Find(&project).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to fetch Project", nil, err.Error())
		return err
	}

	var total int64
	DB.Model(&model.Project{}).Count(&total)

	projectsData := fiber.Map{
		"data": project,
		"meta": fiber.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
		},
	}

	helper.RespondJSON(ctx, fiber.StatusOK, "Project fetched successfully", projectsData, nil)
	return nil
}

func GetProject(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var project model.Project

	// Start a database transaction to ensure atomicity
	tx := database.DB.Begin()

	// Fetch the project including its associated likes
	if err := database.DB.First(&project, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "Project not found", nil, err.Error())
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

	projectData := fiber.Map{
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
		"expired":   project.Expired,
	}

	// Respond with the project data and like count
	helper.RespondJSON(ctx, fiber.StatusOK, "Success get Project", fiber.Map{
		"project": projectData,
	}, nil)

	return nil
}

func CreateProject(ctx *fiber.Ctx) error {
	project := new(model.Project)

	if err := ctx.BodyParser(project); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil, err.Error())
		return err
	}

	if project.Expired.IsZero() {
		if dateString := ctx.FormValue("expired"); dateString != "" {
			layout := "02-01-2006"
			parsedDate, err := time.Parse(layout, dateString)
			if err != nil {
				helper.RespondJSON(ctx, fiber.StatusBadRequest, "Invalid date format. Use dd-MM-yyyy", nil, err.Error())
				return err
			}
			project.Expired = parsedDate
		}
	}

	if err := database.DB.Create(&project).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to create Project", nil, err.Error())
		return err
	}

	helper.RespondJSON(ctx, fiber.StatusOK, "Successfully created Project", project, nil)
	return nil
}

func UpdateProject(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var project model.Project
	if err := database.DB.First(&project, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "Project not found", nil, err.Error())
		return err
	}

	if err := ctx.BodyParser(&project); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil, err.Error())
		return err
	}
	if err := database.DB.Save(&project).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to update Project", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "Successfully update Project", project, nil)
	return nil
}

func DeleteProject(ctx *fiber.Ctx) error {
	if database.DB == nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Database connection not established", nil, nil)
		return errors.New("database connection not established")
	}
	log.Println("Delete Project function called")
	id := ctx.Params("id")
	log.Println("Delete request received for project ID:", id) // Tambahkan logging

	var project model.Project
	if err := database.DB.First(&project, id).Error; err != nil {
		log.Println("Project not found:", err)
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Project not found", nil, err.Error())
		return err
	}

	if err := database.DB.Delete(&project).Error; err != nil {
		log.Println("Failed to delete project:", err)
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to delete Project", nil, err.Error())
		return err
	}
	log.Println("Project deleted successfully")
	helper.RespondJSON(ctx, fiber.StatusOK, "Project deleted successfully", nil, nil)
	return nil
}
