package routes

import (
	"todo-app/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupTaskRoutes(app *fiber.App, db *gorm.DB) {
	taskController := controllers.NewTaskController(db)

	app.Get("/", taskController.GetTasks)
	app.Post("task/add", taskController.AddTask)
	app.Post("task/delete/:id", taskController.DeleteTask)
	app.Post("task/complete/:id", taskController.CompleteTask)
}
