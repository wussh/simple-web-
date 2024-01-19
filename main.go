package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Todo struct represents a to-do item
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var todos []Todo

func main() {
	// Create an Echo instance
	e := echo.New()

	// Routes
	e.GET("/todos", getTodos)
	e.POST("/todos", addTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	// Start the server
	e.Start(":8080")
}

// Handler functions

func getTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func addTodo(c echo.Context) error {
	// Parse request body
	var todo Todo
	err := c.Bind(&todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Assign an ID and add to the list
	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	return c.JSON(http.StatusCreated, todo)
}

func updateTodo(c echo.Context) error {
	// Get todo ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	// Find the todo by ID
	var updatedTodo Todo
	for i, todo := range todos {
		if todo.ID == id {
			err := c.Bind(&updatedTodo)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
			}

			// Update the todo
			updatedTodo.ID = id
			todos[i] = updatedTodo

			return c.JSON(http.StatusOK, updatedTodo)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
}

func deleteTodo(c echo.Context) error {
	// Get todo ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	// Find the todo by ID and remove it
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
}
