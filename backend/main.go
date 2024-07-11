package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

func main() {
	router := gin.Default()
	router.GET("/", greet)
	router.HEAD("/healthcheck", healthcheck)
	router.GET("/items", itemsHandler)
	router.POST("/items", addItem)
	router.GET("/items/:id", getItemByID)
	router.GET("/items/popular", getMostPopularItem)

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

// Item represents a simple structure with an ID, a Name, and a ViewCount.
type Item struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    ViewCount int    `json:"viewCount"`
}

var inventory = []Item{
    {ID: "1", Name: "Galactic Goggles", ViewCount: 0},
    {ID: "2", Name: "Meteor Muffins", ViewCount: 0},
    {ID: "3", Name: "Alien Antenna Kit", ViewCount: 0},
    {ID: "4", Name: "Starlight Lantern", ViewCount: 0},
    {ID: "5", Name: "Quantum Quill", ViewCount: 0},
}

// Use a mutex to safely increment view counts.
var mutex = &sync.Mutex{}

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

func getItemByID(c *gin.Context) {
    id := c.Param("id")

    for i, item := range inventory {
        if item.ID == id {
            // Increment view count in a goroutine to not block the main thread.
            go incrementViewCount(i)

            c.JSON(http.StatusOK, item)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func incrementViewCount(index int) {
    mutex.Lock()
    defer mutex.Unlock()
    inventory[index].ViewCount++
}

func getMostPopularItem(c *gin.Context) {
    if len(inventory) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Inventory is empty"})
        return
    }

    var mostPopular Item
    maxViews := -1

    for _, item := range inventory {
        if item.ViewCount > maxViews {
            mostPopular = item
            maxViews = item.ViewCount
        }
    }

    c.JSON(http.StatusOK, mostPopular)
}
