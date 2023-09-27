package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Item represents an item in the database
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items = []Item{
	{ID: 1, Name: "Item 1", Price: 10.99},
	{ID: 2, Name: "Item 2", Price: 25.50},
}

func main() {
	router := gin.Default()

	// List all items
	router.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, items)
	})

	// Create a new item
	router.POST("/items", func(c *gin.Context) {
		var newItem Item
		if err := c.BindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate a new ID for the item (simple example)
		newItem.ID = len(items) + 1

		items = append(items, newItem)
		c.JSON(http.StatusCreated, newItem)
	})

	// Get a single item by ID
	router.GET("/items/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		for _, item := range items {
			if item.ID == id {
				c.JSON(http.StatusOK, item)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	})

	// Update an item by ID
	router.PUT("/items/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updatedItem Item
		if err := c.BindJSON(&updatedItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, item := range items {
			if item.ID == id {
				items[i] = updatedItem
				c.JSON(http.StatusOK, updatedItem)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	})

	// Delete an item by ID
	router.DELETE("/items/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		for i, item := range items {
			if item.ID == id {
				items = append(items[:i], items[i+1:]...)
				c.Status(http.StatusNoContent)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	})

	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
