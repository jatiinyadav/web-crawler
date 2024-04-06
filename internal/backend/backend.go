package backend

import (
	"github.com/gofiber/fiber/v2"
)

func ExecuteAPI() {
	app := fiber.New()

	// Enable CORS for all routes
	app.Use(func(c *fiber.Ctx) error {
		// Check the Origin header
		origin := c.Get("Origin")

		// If the request is from http://localhost:5173, allow it
		if origin == "http://localhost:5173" {
			c.Set("Access-Control-Allow-Origin", origin)
			c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			c.Set("Access-Control-Allow-Headers", "Content-Type")
		}

		// Handle preflight requests
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()
	})

	// app.Get("/api/message", func(c *fiber.Ctx) error {
	// 	return c.JSON(allitems)
	// })

	app.Listen(":8080")
}
