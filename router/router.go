package router

import (
	"github.com/ahay12/api-test/middleware"
	"github.com/ahay12/api-test/service"
	"github.com/gofiber/fiber/v2"
)

func Cors() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		ctx.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return ctx.Next()
	}
}

func Make() *fiber.App {
	app := fiber.New()
	app.Use(Cors())
	v1 := app.Group("/api/v1")

	// Public routes
	{
		v1.Get("/project", service.GetProjects)
		v1.Get("/project/:id", service.GetProject)
		v1.Post("/signup", service.CreateUser)
		v1.Post("/login", service.Login)
	}
	// Admin routes
	{

		v1.Post("/project", middleware.AdminMiddleware, service.CreateProject)
		v1.Put("/project/:id", middleware.AdminMiddleware, service.UpdateProject)
		v1.Delete("/project/:id", middleware.AdminMiddleware, service.DeleteProject)
	}
	// Admin routes
	{
		v1.Get("/users", middleware.AdminMiddleware, service.GetUsers)
		v1.Get("/user/:id", middleware.AdminMiddleware, service.GetUser)
		v1.Put("/user/:id", middleware.AdminMiddleware, service.UpdateUser)
		v1.Delete("/user/:id", middleware.AdminMiddleware, service.DeleteUser)
	}

	return app
}
