package main

import (
	"fmt"

	"github.com/anujsinghrawat/event-manager/config"
	"github.com/anujsinghrawat/event-manager/db"
	"github.com/anujsinghrawat/event-manager/handlers"
	"github.com/anujsinghrawat/event-manager/middlewares"
	"github.com/anujsinghrawat/event-manager/repositories"
	"github.com/anujsinghrawat/event-manager/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)
	app := fiber.New(fiber.Config{
		AppName:      "Event Manager",
		ServerHeader: "Event Manager",
	})

	// Return a beautiful HTML page for the root
	app.Get("/", func(c *fiber.Ctx) error {
		htmlContent := `<!DOCTYPE html>
			<html>
			<head>
				<meta charset="UTF-8">
				<title>Event Manager</title>
				<style>
					body { font-family: Arial, sans-serif; background: #f4f4f4; text-align: center; padding: 50px; }
					.container { background: #fff; padding: 20px; margin: auto; width: 80%; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
					h1 { color: #333; }
					p { color: #666; }
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Welcome to Event Manager</h1>
					<p>Your event management portal</p>
				</div>
			</body>
			</html>`
		return c.Type("html").SendString(htmlContent)
	})

	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	authService := services.NewAuthService(authRepository)

	server := app.Group("/api/")
	handlers.NewAuthHandler(server.Group("/auth"), authService)
	privateRoutes := server.Use(middlewares.AuthProctected(db))

	handlers.NewEventHandler(privateRoutes.Group("/events"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/tickets"), ticketRepository)

	app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}
