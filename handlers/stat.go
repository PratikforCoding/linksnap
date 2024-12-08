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

// GetURLStats returns the number of clicks for a short URL
func GetURLStats(c *gin.Context) {
	shortCode := c.Param("code")

	// Get the collection
	collection := database.GetCollection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	var url struct {
		ShortCode string `bson:"shortCode"`
		Clicks int		 `bson:"clicks"`
	}

	// Find the URL in the database
	err := collection.FindOne(ctx, bson.M{"shortCode": shortCode}).Decode(&url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	// Return the number of clicks
	c.JSON(http.StatusOK, gin.H{
		"shortCode": url.ShortCode,
		"clicks": url.Clicks,
	})
}