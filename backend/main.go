package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", greet)
	router.HEAD("/healthcheck", healthcheck)
	router.GET("/items", itemsHandler)
	router.POST("/items", addItem)

	router.Run()
}

func greet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome, Go navigator, to the Anythink cosmic catalog.")
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// Item represents a simple structure with an ID and a Name.
type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// Assuming this global slice acts as our inventory.
var inventory := []Item{
	{ID: "1", Name: "Galactic Goggles"},
	{ID: "2", Name: "Meteor Muffins"},
	{ID: "3", Name: "Alien Antenna Kit"},
	{ID: "4", Name: "Starlight Lantern"},
	{ID: "5", Name: "Quantum Quill"},
}

func itemsHandler(c *gin.Context) {
    // Use Gin's method to return JSON response.
    c.JSON(http.StatusOK, inventory)
}

func addItem(c *gin.Context) {
    // Extract name from the request.
    var newItem struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create a new item with a unique ID.
    newID := strconv.Itoa(len(inventory) + 1) // Simple ID generation strategy.
    item := Item{ID: newID, Name: newItem.Name}

    // Add the new item to the inventory.
    inventory = append(inventory, item)

    // Return the new item.
    c.JSON(http.StatusOK, item)
}
