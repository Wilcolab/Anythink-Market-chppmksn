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

func itemsHandler(w http.ResponseWriter, r *http.Request) {
    items := []Item{
        {ID: "1", Name: "Galactic Goggles"},
        {ID: "2", Name: "Meteor Muffins"},
        {ID: "3", Name: "Alien Antenna Kit"},
        {ID: "4", Name: "Starlight Lantern"},
        {ID: "5", Name: "Quantum Quill"},
    }


    // Set the Content-Type header.
    w.Header().Set("Content-Type", "application/json")

    // Encode the items into JSON and send it as the response.
    err := json.NewEncoder(w).Encode(items)
    if err != nil {
        // Handle the error.
        http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
