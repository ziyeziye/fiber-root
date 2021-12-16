package router

import (
	"fiber-root/app/api/login"
	"fiber-root/app/api/task"
	"fiber-root/app/middleware"
	_ "fiber-root/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func InitSysRouter(app *fiber.App) fiber.Router {
	router := app.Group("")
	router.Get("/docs/*", swagger.Handler)
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Hello, World 👋!",
		})
	})

	router.Post("/login", login.Login)

	api := router.Group("/api", middleware.Jwt())

	taskApi := api.Group("/tasks")
	taskApi.Get("/", task.FindAll)
	taskApi.Post("/", task.Save)
	taskApi.Put("/:id", task.ChangeStatus)
	taskApi.Delete("/:id", task.Remove)
	return router
}
