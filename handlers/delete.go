package handlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/PratikforCoding/linksnap/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteURL deletes a short URL from the database
func DeleteURL(c *gin.Context) {
	// Get the short code from the URL
	shortCode := c.Param("code")
	 
	// Get the collection
	collection := database.GetCollection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	// Delete the URL from the database
	result, err := collection.DeleteOne(ctx, bson.M{"shortCode": shortCode})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the URL"})
		return
	}

	// Check if the URL was deleted
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return 
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Short URL deleted successfully"})
}