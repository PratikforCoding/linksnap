package handlers

import(
	"os"
	"context"
	"time"
	"net/http"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
	"github.com/PratikforCoding/linksnap/models"
	"github.com/PratikforCoding/linksnap/database"

)


// RedirectURL redirects the user to the long URL
func RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	// Get the collection
	collection := database.GetCollection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	// Find the URL in the database
	var url models.URL
	err := collection.FindOne(ctx, bson.M{"shortCode": shortCode}).Decode(&url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	c.Redirect(http.StatusFound, url.LongURL)

	// Update the click count
	_, err = collection.UpdateOne(ctx, bson.M{"shortCode": shortCode}, bson.M{
		"$inc": bson.M{"clicks": 1},
	})
	if err != nil {
		fmt.Println("Error updating click count:", err)
	}
}