package main

import (
	"log"
	"todo-app/models"
	"todo-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Initialize the HTML template engine
	htmlEngine := html.New("./views", ".html")

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		Views: htmlEngine,
	})

	app.Use(logger.New())

	// Initialize SQLite database
	var err error
	db, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the Task model
	if err := models.MigrateTask(db); err != nil {
		panic("failed to migrate database")
	}

	routes.SetupTaskRoutes(app, db)

	// Define the port
	port := ":3000"

	// Start the server
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
