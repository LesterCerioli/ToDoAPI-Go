package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"todo-api/database"
	"todo-api/handlers"
	"todo-api/middleware"
	"todo-api/model"
	"todo-api/storage"

	// CRITICAL: Blank import to register swagger docs with swaggo
	// Replace with your actual module path
	_ "todo-api/docs"

	swagger "github.com/Flussen/swagger-fiber-v3"
	"github.com/gofiber/fiber/v3"
)

// @title Todo API
// @version 1.0.0
// @description A production-ready CRUD Todo API with SQLite, Swagger UI, and graceful shutdown.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url https://github.com/yourusername/todo-api
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:3000
// @BasePath /api/v1
// @schemes http
func main() {
	database.Connect()
	database.AutoMigrate(&model.Todo{})

	app := fiber.New(fiber.Config{
		AppName:      "Todo API v1.0",
		ServerHeader: "Fiber",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(middleware.SetupLogger())

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "healthy",
			"version":   "1.0.0",
			"database":  "connected",
			"timestamp": time.Now().UTC(),
		})
	})

	todoStore := storage.NewSQLiteStore()
	todoHandler := handlers.NewTodoHandler(todoStore)

	api := app.Group("/api/v1")
	todos := api.Group("/todos")
	todos.Post("/", todoHandler.Create)
	todos.Get("/", todoHandler.GetAll)
	todos.Get("/:id", todoHandler.GetByID)
	todos.Patch("/:id", todoHandler.Update)
	todos.Delete("/:id", todoHandler.Delete)

	// Swagger UI — now works because docs package is imported
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Endpoint not found",
		})
	})

	go func() {
		log.Println("🚀 Server starting on http://localhost:3000")
		log.Println("📚 Swagger UI: http://localhost:3000/swagger/")
		if err := app.Listen(":3000", fiber.ListenConfig{
			DisableStartupMessage: true,
		}); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Forced shutdown: %v", err)
	}

	log.Println("✅ Server gracefully stopped")
}
