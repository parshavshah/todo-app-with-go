package controllers

import (
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TaskController struct {
	DB *gorm.DB
}

// NewTaskController initializes a new TaskController
func NewTaskController(db *gorm.DB) *TaskController {
	return &TaskController{DB: db}
}

// getTasks renders the main page with the list of tasks
func (tc *TaskController) GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	tc.DB.Find(&tasks)
	return c.Render("index", fiber.Map{
		"Tasks": tasks,
	})
}

// addTask adds a new task to the database
func (tc *TaskController) AddTask(c *fiber.Ctx) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	if (title == "") || (description == "") {
		return c.Redirect("/")
	}
	tc.DB.Create(&models.Task{Title: title, Description: description})
	return c.Redirect("/")
}

// deleteTask removes a task from the database
func (tc *TaskController) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	tc.DB.Delete(&models.Task{}, id)
	return c.Redirect("/")
}

// completeTask marks a task as completed
func (tc *TaskController) CompleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	tc.DB.First(&task, id)
	task.Completed = true
	tc.DB.Save(&task)
	return c.Redirect("/")
}
