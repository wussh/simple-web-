package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create an Echo instance
	e := echo.New()

	// Define a route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// Start the server
	e.Start(":8080")
}
