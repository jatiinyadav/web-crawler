package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/kr/pretty"
)

type details struct {
	Header string `json:"header"`
	Desc   string `json:"desc"`
	Href   string `json:"href"`
}

func main() {
	allitems := []details{}
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div.three div.row div.column", func(e *colly.HTMLElement) {
		items := details{
			Header: e.ChildText("h2.ui"),
			Desc:   e.ChildText("p"),
			Href:   e.ChildAttr("a", "href"),
		}
		allitems = append(allitems, items)
	})

	c.Visit("http://go-colly.org/")
	pretty.Println(allitems)

	// Marshal the struct array into JSON bytes
	jsonData, err := json.Marshal(allitems)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON data to a file
	err = os.WriteFile("products.json", jsonData, 0000)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data written to products.json successfully")

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

	app.Get("/api/message", func(c *fiber.Ctx) error {
		return c.JSON(allitems)
	})

	err = app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
