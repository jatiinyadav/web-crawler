package backend

import (
	"fmt"

	"web-crawler/backend/data"

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
	// 	var item = []Details{}
	// 	return c.JSON(item)
	// })

	app.Get("/api/weburl", func(c *fiber.Ctx) error {
		inputValue1 := c.Query("url1")
		inputValue2 := c.Query("url2")
		fmt.Println(inputValue1)
		fmt.Println(inputValue2)

		result1 := make(chan []data.Details)
		result2 := make(chan []data.Details)
		defer close(result1)
		defer close(result2)

		go data.FetchData(inputValue1, result1)
		go data.FetchData(inputValue2, result2)

		data1 := <-result1
		data2 := <-result2

		results := []data.Details{}
		results = append(results, data1[0], data2[0])

		return c.JSON(results)
	})

	app.Listen(":8080")
}
